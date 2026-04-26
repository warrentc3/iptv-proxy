//go:build go1.18
// +build go1.18

package validate

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoChoice      = errors.New("one or more values is set but there is no possible value available")
	ErrValueNotOneOf = errors.New("value is not one of the possible choices")
)

// IsOneOf returns a `nil` error if the `value` is one of
// the given `possibilities`. Otherwise, an error is returned,
// wrapping `ErrValueNotOneOf` and listing the `possibilities`.
func IsOneOf[T comparable](value T, possibilities ...T) (err error) {
	for _, possibility := range possibilities {
		if value == possibility {
			return nil
		}
	}

	return fmt.Errorf("%w: %v must be one of %s",
		ErrValueNotOneOf, value, orStrings(possibilities))
}

// AreAllOneOf returns a `nil` error if each of the `values`
// are one of the given `choices`. Otherwise, an error is returned,
// wrapping `ErrValueNotOneOf`, precising which value did not match
// and listing the `possibilities`.
func AreAllOneOf[T comparable](values, choices []T) (err error) {
	if len(values) > 0 && len(choices) == 0 {
		return fmt.Errorf("%w", ErrNoChoice)
	}

	set := make(map[T]struct{}, len(choices))
	for _, choice := range choices {
		set[choice] = struct{}{}
	}

	for _, value := range values {
		_, ok := set[value]
		if !ok {
			choiceStrings := make([]string, len(choices))
			for i := range choices {
				choiceStrings[i] = fmt.Sprintf("%v", choices[i])
			}
			return fmt.Errorf("%w: value %v, choices available are %s",
				ErrValueNotOneOf, value, strings.Join(choiceStrings, ", "))
		}
	}

	return nil
}

// AreAllOneOfCaseInsensitive returns a `nil` error if each of the `values`
// are one of the given `choices` in a case insensitive manner. Otherwise,
// an error is returned, wrapping `ErrValueNotOneOf`, precising which value
// did not match and listing the `possibilities`.
func AreAllOneOfCaseInsensitive(values, choices []string) (err error) {
	if len(values) > 0 && len(choices) == 0 {
		return fmt.Errorf("%w", ErrNoChoice)
	}

	set := make(map[string]struct{}, len(choices))
	for _, choice := range choices {
		lowercaseChoice := strings.ToLower(choice)
		set[lowercaseChoice] = struct{}{}
	}

	for _, value := range values {
		lowercaseValue := strings.ToLower(value)
		_, ok := set[lowercaseValue]
		if !ok {
			choiceStrings := make([]string, len(choices))
			for i := range choices {
				choiceStrings[i] = fmt.Sprintf("%v", choices[i])
			}
			return fmt.Errorf("%w: value %v, choices available are %s",
				ErrValueNotOneOf, value, strings.Join(choiceStrings, ", "))
		}
	}

	return nil
}
