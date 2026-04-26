//go:build go1.18
// +build go1.18

package validate

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var ErrValueOutOfBounds = errors.New("value is out of bounds")

// NumberBetween returns a `nil` error if the given `n` is between
// the given `min` and `max` values. Otherwise, an error is returned,
// wrapping `ErrValueOutOfBounds` and describing details on the mismatch.
func NumberBetween[T constraints.Ordered](n, min, max T) (err error) {
	if n < min || n > max {
		return fmt.Errorf("%w: %v must be between %v and %v included",
			ErrValueOutOfBounds, n, min, max)
	}
	return nil
}
