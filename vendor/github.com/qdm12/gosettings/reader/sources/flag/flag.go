package flag

import (
	"strings"
)

// Source implements a CLI flag settings source.
// Note all keys are transformed using its KeyTransform
// method.
type Source struct {
	keyToValue map[string]string
}

// New creates a new flags source from OS arguments.
// WARNING: flags are not typed therefore DO NOT place
// a short boolean flag before a command word, for example
// do not use: ./program --enabled command
// You can however safely use: ./program --enabled=true command
// Boolean short flags without an equal sign or a value after
// have their value set to "true".
// All flag keys read are eventually transformed using
// the KeyTransform method.
func New(osArgs []string) (source *Source) {
	source = &Source{
		keyToValue: make(map[string]string, len(osArgs)),
	}

	if len(osArgs) > 0 {
		// This should always be the case in production
		osArgs = osArgs[1:] // remove the program name
	}

	var key, value string
	for len(osArgs) > 0 {
		key, value, osArgs = parseOne(osArgs)
		if key == "" {
			continue
		}
		key = source.KeyTransform(key)
		source.keyToValue[key] = value
	}

	return source
}

func parseOne(osArgs []string) (key, value string,
	nextOsArgs []string) {
	if len(osArgs) == 0 { // this should not happen
		return "", "", osArgs
	}

	osArg := osArgs[0]
	osArgs = osArgs[1:]

	if !isFlag(osArg) {
		return "", "", osArgs
	}

	osArg = strings.TrimLeft(osArg, "-")

	equalIndex := strings.IndexRune(osArg, '=')
	if equalIndex > 0 {
		key = osArg[:equalIndex]
		value = osArg[equalIndex+1:]
		return key, value, osArgs
	}

	key = osArg

	// Special case for true boolean flags
	valueIsTrue := len(osArgs) == 0 ||
		isFlag(osArgs[0])
	if valueIsTrue {
		value = "true"
		return key, value, osArgs
	}

	value = osArgs[0]
	osArgs = osArgs[1:]
	return key, value, osArgs
}

func isFlag(osArg string) (ok bool) {
	return len(osArg) >= 2 &&
		osArg[0] == '-' &&
		osArg != "-" &&
		osArg != "--" &&
		!strings.HasPrefix(osArg, "---") &&
		!strings.HasPrefix(osArg, "-=") &&
		!strings.HasPrefix(osArg, "--=")
}

func (f *Source) String() string {
	return "flag"
}

// Get returns the value of the flag corresponding
// to the given key, and a boolean `isSet` to
// indicate if it is set or not.
func (f *Source) Get(key string) (value string, isSet bool) {
	value, isSet = f.keyToValue[key]
	return value, isSet
}

// KeyTransform transforms a generic key to a flag
// key. It notably:
// - Changes all characters to be lowercase
// - Replaces all underscores and spaces with dashes.
func (f *Source) KeyTransform(key string) (newKey string) {
	newKey = strings.ToLower(key)
	newKey = strings.ReplaceAll(newKey, "_", "-")
	newKey = strings.ReplaceAll(newKey, " ", "-")
	return newKey
}
