package parse

import (
	"github.com/qdm12/gosettings"
)

type settings struct {
	trimLineEndings     *bool
	trimSpace           *bool
	trimQuotes          *bool
	forceLowercase      *bool
	acceptEmpty         *bool
	currentKey          string
	deprecatedKeys      []string
	handleDeprecatedKey func(source, deprecateKey, currentKey string)
}

func settingsFromOptions(options []Option) (s settings) {
	for _, option := range options {
		option(&s)
	}
	s.setDefaults()
	return s
}

func (s *settings) setDefaults() {
	s.trimLineEndings = gosettings.DefaultPointer(s.trimLineEndings, true)
	s.trimSpace = gosettings.DefaultPointer(s.trimSpace, true)
	s.trimQuotes = gosettings.DefaultPointer(s.trimQuotes, true)
	s.forceLowercase = gosettings.DefaultPointer(s.forceLowercase, true)
	s.acceptEmpty = gosettings.DefaultPointer(s.acceptEmpty, false)
	if s.handleDeprecatedKey == nil {
		s.handleDeprecatedKey = func(source, deprecateKey, currentKey string) {}
	}
}
