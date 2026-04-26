package reader

import (
	"github.com/qdm12/gosettings/internal/parse"
)

// Get returns a value found at the given key as a string pointer.
//
// The value is returned as `nil` if:
//   - the key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set and its corresponding value is empty.
//
// Otherwise, the value may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) Get(key string, options ...Option) (value *string) {
	parseOptions := r.makeParseOptions(options)
	return parse.Get(r.sources, key, parseOptions...)
}

// String returns a string from the value found at the given key,
// which may be modified depending on the parse default settings
// and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the key is not set, the empty string is returned.
func (r *Reader) String(key string, options ...Option) (value string) {
	parseOptions := r.makeParseOptions(options)
	return parse.String(r.sources, key, parseOptions...)
}

// CSV returns a slice of strings from a comma separated value
// found at the given key.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSV(key string, options ...Option) (values []string) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSV(r.sources, key, parseOptions...)
}

// Int returns an `int` from the value found at the given key.
// If the value is not a valid integer string, an error is
// returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int(key string, options ...Option) (n int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int(r.sources, key, parseOptions...)
}

// Int8 returns an `int8` from the value found at the given key.
// If the value is not an integer string between -128 and 127,
// an error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int8(key string, options ...Option) (n int8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int8(r.sources, key, parseOptions...)
}

// Int16 returns an `int16` from the value found at the given key.
// If the value is not an integer string between -32768 and 32767,
// an error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int16(key string, options ...Option) (n int16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int16(r.sources, key, parseOptions...)
}

// Int32 returns an `int32` from the value found at the given key.
// If the value is not an integer string between -2147483648 and
// 2147483647, an error is returned with the source and key in
// its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int32(key string, options ...Option) (n int32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int32(r.sources, key, parseOptions...)
}

// Int64 returns an `int64` from the value found at the given key.
// If the value is not an integer string between -2^63 and 2^63 - 1,
// an error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int64(key string, options ...Option) (n int64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int64(r.sources, key, parseOptions...)
}

// Uint returns an `uint` from the value found at the given key.
// If the value is not a valid unsigned integer string, an error is
// returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint(key string, options ...Option) (n uint, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint(r.sources, key, parseOptions...)
}

// Uint8 returns an `uint8` from the value found at the given key.
// If the value is not an integer string between 0 and 255,
// an error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint8(key string, options ...Option) (n uint8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint8(r.sources, key, parseOptions...)
}

// Uint16 returns an `uint16` from the value found at the given key.
// If the value is not an integer string between 0 and 65535, an
// error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint16(key string, options ...Option) (n uint16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint16(r.sources, key, parseOptions...)
}

// Uint32 returns an `uint32` from the value found at the given key.
// If the value is not an integer string between 0 and 4294967295, an
// error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint32(key string, options ...Option) (n uint32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint32(r.sources, key, parseOptions...)
}

// Uint64 returns an `uint64` from the value found at the given
// key. If the value is not a valid integer string bigger than
// 0, an error is returned with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint64(key string, options ...Option) (n uint64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint64(r.sources, key, parseOptions...)
}

// Float32 returns a `float32` from the value found at the given key.
// If the value is not a valid float64 string, an error is returned
// with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Float32(key string, options ...Option) (f float32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Float32(r.sources, key, parseOptions...)
}

// Float64 returns a `float64` from the value found at the given key.
// If the value is not a valid float64 string, an error is returned
// with the source and key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Float64(key string, options ...Option) (f float64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Float64(r.sources, key, parseOptions...)
}

// BoolPtr returns a pointer to a `bool` from the value found at the given key.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the source and key in its message.
func (r *Reader) BoolPtr(key string, options ...Option) (boolPtr *bool, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.BoolPtr(r.sources, key, parseOptions...)
}

// IntPtr returns a pointer to an `int` from the value found at the given key.
// If the value is not a valid integer string, an error is returned
// with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) IntPtr(key string, options ...Option) (intPtr *int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.IntPtr(r.sources, key, parseOptions...)
}

// Int8Ptr returns a pointer to an `int8` from the value found
// at the given key.
// If the value is not an integer string between -128 to 127,
// an error is returned with the key kind and name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func (r *Reader) Int8Ptr(key string, options ...Option) (
	pointer *int8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int8Ptr(r.sources, key, parseOptions...)
}

// Int16Ptr returns a pointer to an `int16` from the value found
// at the given key.
// If the value is not an integer string between -32768 to 32767,
// an error is returned with the key kind and name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func (r *Reader) Int16Ptr(key string, options ...Option) (
	pointer *int16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int16Ptr(r.sources, key, parseOptions...)
}

// Int32Ptr returns a pointer to an `int32` from the value found
// at the given key.
// If the value is not an integer string between -2147483648 to
// 2147483647, an error is returned with the key kind and name
// in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func (r *Reader) Int32Ptr(key string, options ...Option) (
	pointer *int32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int32Ptr(r.sources, key, parseOptions...)
}

// Int64Ptr returns a pointer to an `int64` from the value found
// at the given key.
// If the value is not an integer string between -2^63 to
// 2^63 - 1, an error is returned with the key kind and name
// in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func (r *Reader) Int64Ptr(key string, options ...Option) (
	pointer *int64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int64Ptr(r.sources, key, parseOptions...)
}

// UintPtr returns a pointer to an `uint` from the value found at the given key.
// If the value is not a valid unsigned integer string matching the
// current system architecture, an error is returned with the key name
// and source name in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) UintPtr(key string, options ...Option) (
	pointer *uint, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.UintPtr(r.sources, key, parseOptions...)
}

// Uint8Ptr returns a pointer to an `uint8` from the value found at the given key.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint8Ptr(key string, options ...Option) (uint8Ptr *uint8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint8Ptr(r.sources, key, parseOptions...)
}

// Uint16Ptr returns a pointer to an `uint16` from the value found at the given key.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the key its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint16Ptr(key string, options ...Option) (
	uint16Ptr *uint16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint16Ptr(r.sources, key, parseOptions...)
}

// Uint32Ptr returns a pointer to an `uint32` from the value found at the given key.
// If the value is not a valid integer string between 0 and 4294967295
// an error is returned with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint32Ptr(key string, options ...Option) (
	uint32Ptr *uint32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint32Ptr(r.sources, key, parseOptions...)
}

// Uint64Ptr returns a pointer to an `uint64` from the value found at the given key.
// If the value is not a valid integer string bigger than 0,
// an error is returned with the source and key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint64Ptr(key string, options ...Option) (
	pointer *uint64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint64Ptr(r.sources, key, parseOptions...)
}

// Float32Ptr returns a pointer to a `float32` from the value
// found at the given key. If the value is not a valid float32
// string, an error is returned with the source and key in its
// message. The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Float32Ptr(key string, options ...Option) (
	pointer *float32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Float32Ptr(r.sources, key, parseOptions...)
}

// Float64Ptr returns a pointer to a `float64` from the value
// found at the given key. If the value is not a valid float64
// string, an error is returned with the source and key in its
// message. The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Float64Ptr(key string, options ...Option) (
	pointer *float64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Float64Ptr(r.sources, key, parseOptions...)
}
