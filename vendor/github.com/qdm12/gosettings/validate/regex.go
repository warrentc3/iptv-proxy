package validate

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrValueMismatchRegex = errors.New("value does not match regular expression")

// MatchRegex returns a `nil` error if the given `value` matches
// the given `regex`. Otherwise, an error is returned, wrapping
// `ErrValueMismatchRegex` and describing details on the mismatch.
func MatchRegex(value string, regex *regexp.Regexp) (err error) {
	if regex.MatchString(value) {
		return nil
	}
	return fmt.Errorf("%w: %q does not match regular expression %s",
		ErrValueMismatchRegex, value, regex)
}

// AllMatchRegex returns a `nil` error if all the given `values`
// match the given `regex`. Otherwise, an error is returned, wrapping
// `ErrValueMismatchRegex` and describing details on the mismatch.
func AllMatchRegex(values []string, regex *regexp.Regexp) (err error) {
	for _, value := range values {
		if err = MatchRegex(value, regex); err != nil {
			return err
		}
	}

	return nil
}
