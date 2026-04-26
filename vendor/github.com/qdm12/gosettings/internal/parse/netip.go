package parse

import (
	"net/netip"
)

// NetipAddr returns a netip.Addr from the first value
// found at the given key from the given sources in order.
// If the value is not a valid netip.Addr string, an error is returned
// with the source and key in its message.
// The value is returned as the empty invalid `netip.Addr{}` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func NetipAddr(sources []Source, key string,
	options ...Option) (addr netip.Addr, err error) {
	return GetParse(sources, key, netip.ParseAddr, options...)
}

// NetipAddrPort returns a netip.AddrPort from the first value
// found at the given key from the given sources in order.
// If the value is not a valid netip.AddrPort string, an error is returned
// with the source and key in its message.
// The value is returned as the empty invalid `netip.AddrPort{}` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func NetipAddrPort(sources []Source, key string,
	options ...Option) (addrPort netip.AddrPort, err error) {
	return GetParse(sources, key, netip.ParseAddrPort, options...)
}

// NetipPrefix returns a netip.Prefix from the first value
// found at the given key from the given sources in order.
// If the value is not a valid netip.Prefix string, an error is returned
// with the key name and kind (source) in its message.
// The value is returned as the empty invalid `netip.Prefix{}` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func NetipPrefix(sources []Source, key string,
	options ...Option) (addr netip.Prefix, err error) {
	return GetParse(sources, key, netip.ParsePrefix, options...)
}

// CSVNetipAddresses returns a slice of netip.Addr from the
// first comma separated string value found at the given key
// from the given sources in order, and returns an error if
// any value is not a valid netip.Addr string.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
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
func CSVNetipAddresses(sources []Source, key string,
	options ...Option) (addresses []netip.Addr, err error) {
	return CSVParse(sources, key, netip.ParseAddr, options...)
}

// CSVNetipAddrPorts returns a slice of netip.AddrPort from the
// first comma separated string value found at the given key
// from the given sources in order, and returns an error if
// any value is not a valid netip.AddrPort string.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
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
func CSVNetipAddrPorts(sources []Source, key string,
	options ...Option) (addrPorts []netip.AddrPort, err error) {
	return CSVParse(sources, key, netip.ParseAddrPort, options...)
}

// CSVNetipPrefixes returns a slice of netip.Prefix from
// the first comma separated string value found at the given
// key from the given sources in order, and returns an error
// if any value is not a valid netip.Prefix string.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
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
func CSVNetipPrefixes(sources []Source, key string,
	options ...Option) (prefixes []netip.Prefix, err error) {
	return CSVParse(sources, key, netip.ParsePrefix, options...)
}
