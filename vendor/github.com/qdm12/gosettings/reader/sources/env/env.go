package env

import (
	"strings"
)

// Source implements an environment variables source.
// Note all keys are transformed using its KeyTransform
// method.
type Source struct {
	keyToValue map[string]string
	keyPrefix  string
}

// New creates a new environment variable source
// given the environment as a slice of key-value pairs,
// which can generally be obtained with os.Environ().
// Environ values with no '=' sign are ignored.
// All environment variable keys read are eventually
// transformed using the KeyTransform method.
func New(settings Settings) (source *Source) {
	settings.setDefaults()
	source = &Source{
		keyToValue: make(map[string]string, len(settings.Environ)),
		keyPrefix:  settings.KeyPrefix,
	}

	for _, keyValue := range settings.Environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			continue
		}

		key := source.KeyTransform(parts[0])
		value := parts[1]
		source.keyToValue[key] = value
	}

	return source
}

func (s *Source) String() string {
	return "environment variable"
}

// Get returns the value of the environment variable
// found at the given key, and a boolean `isSet` to
// indicate if it is set or not.
func (s *Source) Get(key string) (value string, isSet bool) {
	value, isSet = s.keyToValue[key]
	return value, isSet
}

// KeyTransform transforms a generic key to an environment
// variable key. It notably:
// - Changes all characters to be uppercase
// - Replaces all dashes with underscores.
// - Prefixes the key with the KeyPrefix field, without modifying the prefix.
func (s *Source) KeyTransform(key string) (newKey string) {
	newKey = strings.ToUpper(key)
	newKey = strings.ReplaceAll(newKey, "-", "_")
	newKey = s.keyPrefix + newKey
	return newKey
}
