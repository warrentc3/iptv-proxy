package parse

import (
	"fmt"
)

// CSVParse returns a slice of type T from the first comma separated
// string value found at the given key in the given sources in order.
// Each comma separated values is parsed using the provided
// `parse` function.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVParse[T any](sources []Source, key string,
	parse ParseFunc[T], options ...Option) (values []T, err error) {
	stringValues, sourceName := csv(sources, key, options...)
	if stringValues == nil {
		return nil, nil
	}

	values = make([]T, len(stringValues))
	for i, stringValue := range stringValues {
		values[i], err = parse(stringValue)
		if err != nil {
			return nil, fmt.Errorf("%s %s: %w", sourceName, key, err)
		}
	}

	return values, nil
}

// CSVInt returns a slice of int from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid int string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVInt(sources []Source, key string,
	options ...Option) (values []int, err error) {
	return CSVParse(sources, key, parseInt, options...)
}

// CSVInt8 returns a slice of int8 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid int8 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVInt8(sources []Source, key string,
	options ...Option) (values []int8, err error) {
	return CSVParse(sources, key, parseInt8, options...)
}

// CSVInt16 returns a slice of int16 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid int16 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVInt16(sources []Source, key string,
	options ...Option) (values []int16, err error) {
	return CSVParse(sources, key, parseInt16, options...)
}

// CSVInt32 returns a slice of int32 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid int32 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVInt32(sources []Source, key string,
	options ...Option) (values []int32, err error) {
	return CSVParse(sources, key, parseInt32, options...)
}

// CSVInt64 returns a slice of int64 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid int64 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVInt64(sources []Source, key string,
	options ...Option) (values []int64, err error) {
	return CSVParse(sources, key, parseInt64, options...)
}

// CSVUint returns a slice of uint from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid uint string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVUint(sources []Source, key string,
	options ...Option) (values []uint, err error) {
	return CSVParse(sources, key, parseUint, options...)
}

// CSVUint8 returns a slice of uint8 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVUint8(sources []Source, key string,
	options ...Option) (values []uint8, err error) {
	return CSVParse(sources, key, parseUint8, options...)
}

// CSVUint16 returns a slice of uint8 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVUint16(sources []Source, key string,
	options ...Option) (values []uint16, err error) {
	return CSVParse(sources, key, parseUint16, options...)
}

// CSVUint32 returns a slice of uint32 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid uint32 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVUint32(sources []Source, key string,
	options ...Option) (values []uint32, err error) {
	return CSVParse(sources, key, parseUint32, options...)
}

// CSVUint64 returns a slice of uint64 from the first comma separated
// string value found at the given key in the given sources in order.
// It returns an error if any value is not a valid uint64 string.
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and the corresponding value is empty.
func CSVUint64(sources []Source, key string,
	options ...Option) (values []uint64, err error) {
	return CSVParse(sources, key, parseUint64, options...)
}
