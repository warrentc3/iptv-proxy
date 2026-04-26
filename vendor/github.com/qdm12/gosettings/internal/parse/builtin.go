package parse

import (
	"fmt"
	"strings"
)

// String returns the first string value found at the
// given key from the given sources in order.
// The value may be modified depending on the parse default settings
// and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the key is not set in any of the sources, the empty string is
// returned.
func String(sources []Source, key string,
	options ...Option) (value string) {
	s, _ := get(sources, key, options...)
	if s == nil {
		return ""
	}
	return *s
}

// CSV returns a slice of strings from the first comma separated
// value found from the given sources in order.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func CSV(sources []Source, key string,
	options ...Option) (values []string) {
	csv, _ := csv(sources, key, options...)
	return csv
}

func csv(sources []Source, key string,
	options ...Option) (values []string, sourceName string) {
	csv, sourceName := get(sources, key, options...)
	if csv == nil {
		return nil, sourceName
	}
	return strings.Split(*csv, ","), sourceName
}

// GetParse parses the first value found at the given key
// from the given sources in order, using the given parse
// function, and returns the typed parsed value with an
// eventual error containing the key name and the source name
// in its message.
//
// The value is returned as the empty `T` value if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in a source and its corresponding value is empty.
//
// Note the value from the map may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func GetParse[T any](sources []Source, key string, //nolint:ireturn
	parse ParseFunc[T], options ...Option) (value T, err error) {
	s, sourceKind := get(sources, key, options...)
	if s == nil {
		return value, nil
	}

	value, err = parse(*s)
	if err != nil {
		return value, fmt.Errorf("%s %s: %w", sourceKind, key, err)
	}

	return value, nil
}

// GetParsePtr parses the first value found at the given key from
// the given sources in order, using the given parse function, and
// returns the typed pointer to the parsed value with an eventual
// error which contains the key name and source name in its message.
//
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in the mapping and its corresponding value is empty.
//
// Note the value from the map may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func GetParsePtr[T any](sources []Source, key string,
	parse ParseFunc[T], options ...Option) (value *T, err error) {
	s, sourceKind := get(sources, key, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value = new(T)
	*value, err = parse(*s)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %w", sourceKind, key, err)
	}

	return value, nil
}

// Int returns an `int` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int(sources []Source, key string,
	options ...Option) (n int, err error) {
	return GetParse(sources, key, parseInt, options...)
}

// Int8 returns an `int8` from the first value found at the given
// key in the given sources in order.
// If the value is not an integer string between -128 and 127, an
// error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int8(sources []Source, key string,
	options ...Option) (n int8, err error) {
	return GetParse(sources, key, parseInt8, options...)
}

// Int16 returns an `int16` from the first value found at the given
// key in the given sources in order.
// If the value is not an integer string between -32768 and 32767, an
// error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int16(sources []Source, key string,
	options ...Option) (n int16, err error) {
	return GetParse(sources, key, parseInt16, options...)
}

// Int32 returns an `int32` from the first value found at the given
// key in the given sources in order.
// If the value is not an integer string between -2147483648 and
// 2147483647, an error is returned with the key name and the
// source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int32(sources []Source, key string,
	options ...Option) (n int32, err error) {
	return GetParse(sources, key, parseInt32, options...)
}

// Int64 returns an `int64` from the first value found at the given
// key in the given sources in order.
// If the value is not an integer string between –2^63 and 2^63 - 1,
// an error is returned with the key name and the source name in
// its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int64(sources []Source, key string,
	options ...Option) (n int64, err error) {
	return GetParse(sources, key, parseInt64, options...)
}

// Uint returns an `uint` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid unsigned integer string matching the
// current system architecture, an error is returned with the key
// name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Uint(sources []Source, key string,
	options ...Option) (n uint, err error) {
	return GetParse(sources, key, parseUint, options...)
}

// Uint8 returns an `uint8` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string between 0 and 255, an
// error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Uint8(sources []Source, key string,
	options ...Option) (n uint8, err error) {
	return GetParse(sources, key, parseUint8, options...)
}

// Uint16 returns an `uint16` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string between 0 and 65535, an
// error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Uint16(sources []Source, key string,
	options ...Option) (n uint16, err error) {
	return GetParse(sources, key, parseUint16, options...)
}

// Uint32 returns an `uint32` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string between 0 and 4294967295, an
// error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Uint32(sources []Source, key string,
	options ...Option) (n uint32, err error) {
	return GetParse(sources, key, parseUint32, options...)
}

// Uint64 returns an `uint64` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string between 0 and 18446744073709551615,
// an error is returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Uint64(sources []Source, key string,
	options ...Option) (n uint64, err error) {
	return GetParse(sources, key, parseUint64, options...)
}

// Float32 returns a `float32` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid float32 string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float32(sources []Source, key string,
	options ...Option) (f float32, err error) {
	return GetParse(sources, key, parseFloat32, options...)
}

// Float64 returns a `float64` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid float64 string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float64(sources []Source, key string,
	options ...Option) (f float64, err error) {
	return GetParse(sources, key, parseFloat64, options...)
}

// BoolPtr returns a pointer to a `bool` from the first value found
// at the given key in the given sources in order.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the key name and source name in its message.
func BoolPtr(sources []Source, key string, options ...Option) (
	boolPtr *bool, err error) {
	return GetParse(sources, key, parseBool, options...)
}

// IntPtr returns a pointer to an `int` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string, an error is returned
// with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func IntPtr(sources []Source, key string, options ...Option) (
	intPtr *int, err error) {
	return GetParsePtr(sources, key, parseInt, options...)
}

// Int8Ptr returns a pointer to an `int8` from the first value found
// at the given key in the given sources in order.
// If the value is not an integer string between -128 to 127, an error
// is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Int8Ptr(sources []Source, key string, options ...Option) (
	pointer *int8, err error) {
	return GetParsePtr(sources, key, parseInt8, options...)
}

// Int16Ptr returns a pointer to an `int8` from the first value found
// at the given key in the given sources in order.
// If the value is not an integer string between -32768 to 32767, an
// error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Int16Ptr(sources []Source, key string, options ...Option) (
	pointer *int16, err error) {
	return GetParsePtr(sources, key, parseInt16, options...)
}

// Int32Ptr returns a pointer to an `int32` from the first value found
// at the given key in the given sources in order.
// If the value is not an integer string between -2147483648 to 2147483647,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Int32Ptr(sources []Source, key string, options ...Option) (
	pointer *int32, err error) {
	return GetParsePtr(sources, key, parseInt32, options...)
}

// Int64Ptr returns a pointer to an `int64` from the first value found
// at the given key in the given sources in order.
// If the value is not an integer string between -2^63 to 2^63 - 1,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Int64Ptr(sources []Source, key string, options ...Option) (
	pointer *int64, err error) {
	return GetParsePtr(sources, key, parseInt64, options...)
}

// UintPtr returns a pointer to an `uint` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid unsigned integer string matching the
// current system architecture, an error is returned with the key name
// and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func UintPtr(sources []Source, key string, options ...Option) (
	pointer *uint, err error) {
	return GetParsePtr(sources, key, parseUint, options...)
}

// Uint8Ptr returns a pointer to an `uint8` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint8Ptr(sources []Source, key string, options ...Option) (
	uint8Ptr *uint8, err error) {
	return GetParsePtr(sources, key, parseUint8, options...)
}

// Uint16Ptr returns a pointer to an `uint16` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint16Ptr(sources []Source, key string, options ...Option) (
	uint16Ptr *uint16, err error) {
	return GetParsePtr(sources, key, parseUint16, options...)
}

// Uint32Ptr returns a pointer to an `uint32` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 4294967295,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint32Ptr(sources []Source, key string, options ...Option) (
	uint32Ptr *uint32, err error) {
	return GetParsePtr(sources, key, parseUint32, options...)
}

// Uint64Ptr returns a pointer to an `uint64` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string bigger than 0,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint64Ptr(sources []Source, key string, options ...Option) (
	pointer *uint64, err error) {
	return GetParsePtr(sources, key, parseUint64, options...)
}

// Float32 returns a pointer to a `float32` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid float32 string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float32Ptr(sources []Source, key string,
	options ...Option) (pointer *float32, err error) {
	return GetParsePtr(sources, key, parseFloat32, options...)
}

// Float64 returns a `float64` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid float64 string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float64Ptr(sources []Source, key string,
	options ...Option) (pointer *float64, err error) {
	return GetParsePtr(sources, key, parseFloat64, options...)
}
