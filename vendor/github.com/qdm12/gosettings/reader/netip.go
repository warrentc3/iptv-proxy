package reader

import (
	"net/netip"

	"github.com/qdm12/gosettings/internal/parse"
)

// NetipAddr returns a netip.Addr from the value found at the given key.
// If the value is not a valid netip.Addr string, an error is returned
// with the source and key in its message.
// The value is returned as the empty invalid `netip.Addr{}` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) NetipAddr(key string, options ...Option) (
	addr netip.Addr, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.NetipAddr(r.sources, key, parseOptions...)
}

// NetipAddrPort returns a netip.AddrPort from the value found at the given key.
// If the value is not a valid netip.AddrPort string, an error is returned
// with the source and key in its message.
// The value is returned as the empty invalid `netip.AddrPort{}` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) NetipAddrPort(key string, options ...Option) (
	addr netip.AddrPort, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.NetipAddrPort(r.sources, key, parseOptions...)
}

// NetipPrefix returns a netip.Prefix from the value found at the given key.
// If the value is not a valid netip.Prefix string, an error is returned
// with the source and key in its message.
// The value is returned as the empty invalid `netip.Prefix{}` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) NetipPrefix(key string, options ...Option) (
	addr netip.Prefix, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.NetipPrefix(r.sources, key, parseOptions...)
}

// CSVNetipAddresses returns a slice of netip.Addr from a comma separated value
// found at the given key and returns an error if any value
// is not a valid netip.Addr string.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) CSVNetipAddresses(key string, options ...Option) (
	addresses []netip.Addr, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVNetipAddresses(r.sources, key, parseOptions...)
}

// CSVNetipAddrPorts returns a slice of netip.AddrPort from a
// comma separated value found at the given key and returns an
// error if any value is not a valid netip.AddrPort string.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) CSVNetipAddrPorts(key string, options ...Option) (
	addrPorts []netip.AddrPort, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVNetipAddrPorts(r.sources, key, parseOptions...)
}

// CSVNetipPrefixes returns a slice of netip.Prefix from a comma separated value
// found at the given key and returns an error if any value
// is not a valid netip.Prefix string.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) CSVNetipPrefixes(key string, options ...Option) (
	prefixes []netip.Prefix, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVNetipPrefixes(r.sources, key, parseOptions...)
}
