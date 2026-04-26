package parse

import (
	"slices"
	"strings"
)

// Get returns the first value found at the given key from
// the given sources in order, as a string pointer.
//
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in one of the sources and its corresponding value is empty.
//
// Otherwise, the value may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func Get(sources []Source, key string, options ...Option) (value *string) {
	value, _ = get(sources, key, options...)
	return value
}

func get(sources []Source, key string, options ...Option) (
	value *string, sourceKind string) {
	settings := settingsFromOptions(options)

	keysToTry := make([]string, 0, 1+len(settings.deprecatedKeys))
	keysToTry = append(keysToTry, settings.deprecatedKeys...)
	// Note we try the current key last because it might be set
	// in the released program (such as a Docker image), so we want
	// to take the older configuration from the user first.
	keysToTry = append(keysToTry, key)

	var firstKeySet string
	var firstSource Source

	for _, keyToTry := range keysToTry {
		for _, sourceToTry := range sources {
			transformedKeyToTry := sourceToTry.KeyTransform(keyToTry)
			stringValue, isSet := sourceToTry.Get(transformedKeyToTry)
			if !isSet || (!*settings.acceptEmpty && stringValue == "") {
				continue
			}
			firstKeySet = transformedKeyToTry
			firstSource = sourceToTry
			value = new(string)
			*value = stringValue
			break
		}
		if firstKeySet != "" {
			break
		}
	}

	if firstKeySet == "" { // All keys are unset for all sources
		return nil, ""
	}

	key = firstSource.KeyTransform(key)
	sourceKind = firstSource.String()
	if settings.currentKey != "" { // all keys are retro-compatible keys
		currentKey := firstSource.KeyTransform(settings.currentKey)
		settings.handleDeprecatedKey(sourceKind, firstKeySet, currentKey)
	} else if firstKeySet != key {
		settings.handleDeprecatedKey(sourceKind, firstKeySet, key)
	}

	*value = postProcessValue(*value, settings)
	return value, sourceKind
}

func postProcessValue(value string, settings settings) string {
	if *settings.forceLowercase {
		value = strings.ToLower(value)
	}

	cutSet := map[string]struct{}{}
	if *settings.trimSpace {
		// Only latin charset
		spaceCharacters := []rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}
		for _, r := range spaceCharacters {
			cutSet[string(r)] = struct{}{}
		}
	}

	if *settings.trimLineEndings {
		cutSet["\r"] = struct{}{}
		cutSet["\n"] = struct{}{}
	}

	cutSetString := ""
	for s := range cutSet {
		cutSetString += s
	}

	quotes := []rune{'\'', '"', '`'}
	for {
		value = strings.Trim(value, cutSetString)
		const minCharactersForQuotes = 2
		if *settings.trimQuotes &&
			len(value) >= minCharactersForQuotes &&
			slices.Contains(quotes, rune(value[0])) &&
			slices.Contains(quotes, rune(value[len(value)-1])) {
			value = value[1 : len(value)-1] // remove quotes
			continue                        // re-trim cut set
		}
		break
	}
	return value
}
