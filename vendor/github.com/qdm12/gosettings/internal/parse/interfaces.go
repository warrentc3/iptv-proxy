package parse

// Source is a named key-value source.
type Source interface {
	// String can return for example 'environment variable' or 'flag'
	String() string
	// Get returns the value of the key and whether it is set.
	Get(key string) (value string, isSet bool)
	// KeyTransform transforms a standardized key to a key specific to
	// the source. For example SERVER_ADDRESS becomes server-address for
	// the flags source.
	KeyTransform(key string) string
}
