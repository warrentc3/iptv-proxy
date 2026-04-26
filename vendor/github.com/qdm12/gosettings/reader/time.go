package reader

import (
	"time"

	"github.com/qdm12/gosettings/internal/parse"
)

// DurationPtr returns a pointer to a `time.Duration`
// from the value found at the given key.
// If the value is not a valid time.Duration string, an error
// is returned with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) DurationPtr(key string, options ...Option) (
	durationPtr *time.Duration, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.DurationPtr(r.sources, key, parseOptions...)
}

// Duration returns a `time.Duration` from the value found at
// the given key.
// If the value is not a valid time.Duration string, an error
// is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Duration(key string, options ...Option) (
	duration time.Duration, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Duration(r.sources, key, parseOptions...)
}
