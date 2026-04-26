package parse

import (
	"time"
)

// DurationPtr returns a pointer to a `time.Duration`
// from the first value found at the given key from the
// given sources in order.
// If the value is not a valid time.Duration string, an error
// is returned with the source name and key in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func DurationPtr(sources []Source, key string,
	options ...Option) (durationPtr *time.Duration, err error) {
	return GetParsePtr(sources, key, time.ParseDuration, options...)
}

// Duration returns a `time.Duration` parsed from the first
// value found at the given key from the given sources in order.
// If the value is not a valid time.Duration string, an error
// is returned with the source name and key in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Duration(sources []Source, key string,
	options ...Option) (duration time.Duration, err error) {
	return GetParse(sources, key, time.ParseDuration, options...)
}
