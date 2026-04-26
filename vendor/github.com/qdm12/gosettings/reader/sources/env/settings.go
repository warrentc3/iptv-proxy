package env

import (
	"os"

	"github.com/qdm12/gosettings"
)

// Settings contains settings for the environment variable source.
type Settings struct {
	// Environ is a slice of "key=value" pairs which defaults
	// to the returned value of os.Environ().
	// Element without a '=' sign are ignored.
	// All environment variable keys read are eventually
	// transformed using the KeyTransform method.
	// It defaults to the empty string.
	Environ []string
	// KeyPrefix is a prefix to add to all keys read from
	// the environment. It is usually the program name to avoid
	// conflict with other programs in the environment.
	// It defaults to the empty string.
	KeyPrefix string
}

func (s *Settings) setDefaults() {
	s.Environ = gosettings.DefaultSlice(s.Environ, os.Environ())
}
