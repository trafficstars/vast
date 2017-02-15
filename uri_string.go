package vast

import "strings"

// URIString is a string that allows for stripping of whitespace when Unmarshalled
type URIString string

// MarshalText implements the encoding.TextMarshaler interface.
func (s URIString) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *URIString) UnmarshalText(data []byte) error {
	*s = URIString(strings.TrimSpace(string(data)))
	return nil
}
