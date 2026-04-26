package parse

// Option is an option to modify the behavior of the
// underlying `get` function which is called by all the
// other functions.
type Option func(s *settings)

// ForceLowercase forces the string values read from any
// source given to be lowercased or not, depending on the
// `lowercase` argument given.
func ForceLowercase(lowercase bool) Option {
	return func(s *settings) {
		s.forceLowercase = &lowercase
	}
}

// AcceptEmpty, if set to true, makes the code distinguish
// between unset keys and empty values from a given source.
// By default, the code does not distinguish between the two cases.
func AcceptEmpty(accept bool) Option {
	return func(s *settings) {
		s.acceptEmpty = &accept
	}
}

// RetroKeys specifies a list of keys that are deprecated and
// replaced by the current key.
// The oldest deprecated key should be placed first
// in the list of deprecated keys, so it gets checked first, which
// is especially important if for example default variable values
// are set in the program or operating system matching more recent keys.
// The `handleDeprecatedKey` function is called when a deprecated
// key is used, with the source name, the deprecated key and
// the current key as arguments.
func RetroKeys(handleDeprecatedKey func(source, deprecateKey, currentKey string),
	deprecatedKeys ...string) Option {
	return func(s *settings) {
		s.deprecatedKeys = deprecatedKeys
		s.handleDeprecatedKey = handleDeprecatedKey
	}
}

// IsRetro indicates that all the keys given are retro-compatible keys,
// and that the currentKey given should be used instead.
func IsRetro(handleDeprecatedKey func(source, deprecateKey, currentKey string),
	currentKey string) Option {
	return func(s *settings) {
		s.currentKey = currentKey
		s.handleDeprecatedKey = handleDeprecatedKey
	}
}
