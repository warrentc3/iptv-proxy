package reader

import "github.com/qdm12/gosettings/internal/parse"

// CSVInt returns a slice of int from a comma separated value
// found at the given key and returns an error if any value
// is not a valid int string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVInt(key string, options ...Option) (values []int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt(r.sources, key, parseOptions...)
}

// CSVInt8 returns a slice of int8 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid int8 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVInt8(key string, options ...Option) (values []int8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt8(r.sources, key, parseOptions...)
}

// CSVInt16 returns a slice of int16 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid int16 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVInt16(key string, options ...Option) (values []int16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt16(r.sources, key, parseOptions...)
}

// CSVInt32 returns a slice of int32 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid int32 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVInt32(key string, options ...Option) (values []int32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt32(r.sources, key, parseOptions...)
}

// CSVInt64 returns a slice of int64 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid int64 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVInt64(key string, options ...Option) (values []int64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt64(r.sources, key, parseOptions...)
}

// CSVUint returns a slice of uint from a comma separated value
// found at the given key and returns an error if any value
// is not a valid uint string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVUint(key string, options ...Option) (values []uint, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint(r.sources, key, parseOptions...)
}

// CSVUint8 returns a slice of uint8 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVUint8(key string, options ...Option) (values []uint8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint8(r.sources, key, parseOptions...)
}

// CSVUint16 returns a slice of uint16 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid uint16 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVUint16(key string, options ...Option) (values []uint16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint16(r.sources, key, parseOptions...)
}

// CSVUint32 returns a slice of uint32 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid uint32 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVUint32(key string, options ...Option) (values []uint32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint32(r.sources, key, parseOptions...)
}

// CSVUint64 returns a slice of uint64 from a comma separated value
// found at the given key and returns an error if any value
// is not a valid uint64 string.
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSVUint64(key string, options ...Option) (values []uint64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint64(r.sources, key, parseOptions...)
}
