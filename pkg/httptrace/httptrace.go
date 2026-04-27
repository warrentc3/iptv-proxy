// v4 release

/*
 * Iptv-Proxy is a project to proxyfie an m3u file and to proxyfie an Xtream iptv service (client API).
 * Copyright (C) 2026  warrentc3
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Package httptrace provides opt-in, env-gated tracing of every inbound and
// outbound HTTP call the proxy makes. Enabled when DEBUG_HTTP is truthy at
// startup; sink defaults to stdout, redirects to DEBUG_HTTP_FILE when set.
//
// The trace exists as an empirical-validation tool: live confirmation that
// what the proxy receives, constructs, and forwards matches expectation.
// Auth headers are logged in plaintext — this is correct behavior for the
// debug-tool framing and the env gate is the safety boundary.
package httptrace

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const bodyMaxBytes int64 = 64 * 1024

var (
	once    sync.Once
	enabled bool
	sink    *log.Logger
)

// Init configures the trace sink. Idempotent: only the first call takes
// effect. Safe to call when disabled (sink becomes a no-op writer).
func Init(on bool, filePath string) {
	once.Do(func() {
		enabled = on
		if !enabled {
			sink = log.New(io.Discard, "", 0)
			return
		}
		var w io.Writer = os.Stdout
		if filePath != "" {
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
			if err != nil {
				log.Printf("[iptv-proxy] DEBUG_HTTP_FILE %q open failed (%v); falling back to stdout", filePath, err)
			} else {
				w = f
			}
		}
		sink = log.New(w, "[HTTP-TRACE] ", log.LstdFlags|log.Lmicroseconds)
		sink.Printf("trace enabled; bodyMaxBytes=%d", bodyMaxBytes)
	})
}

// Enabled reports whether trace was turned on at Init time.
func Enabled() bool { return enabled }

// WrapTransport returns base wrapped with a logging RoundTripper when trace
// is enabled, otherwise returns base unchanged. base may be nil; in that
// case http.DefaultTransport is used as the underlying transport.
func WrapTransport(base http.RoundTripper) http.RoundTripper {
	if !enabled {
		return base
	}
	if base == nil {
		base = http.DefaultTransport
	}
	return &loggingTransport{base: base}
}

type loggingTransport struct {
	base http.RoundTripper
}

func (lt *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	// Capture request body if it fits the buffer-friendly content profile.
	// Reads the FULL body (not LimitReader) so the downstream consumer
	// receives the complete payload — the bodyMaxBytes cap is for log
	// display only, applied in logCapturedBody.
	var reqBody []byte
	if req.Body != nil && shouldLogBody(req.Header.Get("Content-Type"), req.ContentLength) {
		reqBody, _ = io.ReadAll(req.Body)
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewReader(reqBody))
	}

	sink.Printf("OUT --> %s %s", req.Method, req.URL.String())
	for k, vv := range req.Header {
		for _, v := range vv {
			sink.Printf("OUT --> [hdr] %s: %s", k, v)
		}
	}
	if len(reqBody) > 0 {
		logCapturedBody("OUT -->", reqBody)
	} else if req.Body != nil {
		sink.Printf("OUT --> [body present, not buffered]")
	}

	resp, err := lt.base.RoundTrip(req)
	dur := time.Since(start)
	if err != nil {
		sink.Printf("OUT <-- %s %s ERROR after %s: %v", req.Method, req.URL.String(), dur, err)
		return resp, err
	}

	sink.Printf("OUT <-- %s %s %s after %s", req.Method, req.URL.String(), resp.Status, dur)
	for k, vv := range resp.Header {
		for _, v := range vv {
			sink.Printf("OUT <-- [hdr] %s: %s", k, v)
		}
	}

	if shouldLogBody(resp.Header.Get("Content-Type"), resp.ContentLength) {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(body))
		logCapturedBody("OUT <--", body)
	} else {
		sink.Printf("OUT <-- [body content-type=%q content-length=%d, not buffered]",
			resp.Header.Get("Content-Type"), resp.ContentLength)
	}

	return resp, nil
}

// logCapturedBody emits a body line, logging up to bodyMaxBytes of the body
// for display. The full body is preserved by the caller for downstream
// consumers — this function only controls what gets written to the trace.
func logCapturedBody(prefix string, body []byte) {
	if int64(len(body)) > bodyMaxBytes {
		sink.Printf("%s [body %d total bytes, logging first %d]\n%s",
			prefix, len(body), bodyMaxBytes, string(body[:bodyMaxBytes]))
		return
	}
	sink.Printf("%s [body %d bytes]\n%s", prefix, len(body), string(body))
}

// GinMiddleware returns a gin middleware that logs every inbound request and
// its response. No-op when trace is disabled.
func GinMiddleware() gin.HandlerFunc {
	if !enabled {
		return func(c *gin.Context) { c.Next() }
	}
	return func(c *gin.Context) {
		start := time.Now()

		sink.Printf("IN  --> %s %s%s", c.Request.Method, c.Request.URL.Path, queryString(c.Request.URL.RawQuery))
		sink.Printf("IN  --> [client] %s", c.ClientIP())
		for k, vv := range c.Request.Header {
			for _, v := range vv {
				sink.Printf("IN  --> [hdr] %s: %s", k, v)
			}
		}

		if c.Request.Body != nil && shouldLogBody(c.Request.Header.Get("Content-Type"), c.Request.ContentLength) {
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
			logCapturedBody("IN  -->", body)
		}

		c.Next()

		dur := time.Since(start)
		sink.Printf("IN  <-- %s %s %d after %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), dur)
		for k, vv := range c.Writer.Header() {
			for _, v := range vv {
				sink.Printf("IN  <-- [hdr] %s: %s", k, v)
			}
		}
	}
}

// shouldLogBody decides whether to buffer-and-log a request or response body.
// Bodies are skipped for video/audio/octet-stream content (would break or
// exhaust memory on streams) and for any payload exceeding bodyMaxBytes.
// Empty Content-Type is treated conservatively as "do not buffer" because
// a missing type often signals a stream.
func shouldLogBody(contentType string, contentLength int64) bool {
	if contentLength > bodyMaxBytes {
		return false
	}
	ct := strings.ToLower(contentType)
	if ct == "" {
		return false
	}
	if strings.HasPrefix(ct, "video/") || strings.HasPrefix(ct, "audio/") {
		return false
	}
	if strings.Contains(ct, "octet-stream") {
		return false
	}
	if strings.HasPrefix(ct, "application/json") ||
		strings.HasPrefix(ct, "application/xml") ||
		strings.HasPrefix(ct, "application/x-www-form-urlencoded") ||
		strings.HasPrefix(ct, "text/") {
		return true
	}
	return false
}

func queryString(raw string) string {
	if raw == "" {
		return ""
	}
	return "?" + raw
}
