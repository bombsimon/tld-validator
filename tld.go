package tld

import (
	"strings"

	"golang.org/x/net/idna"
)

//go:generate go run cmd/tld-generator/main.go

// TLD represents a top level domain
type TLD string

// String returns the string value for a TLD.
func (t TLD) String() string {
	return strings.ToUpper(string(t))
}

// LowerString returns the string value for a TLD as lowercase.
func (t TLD) LowerString() string {
	return strings.ToLower(string(t))
}

// AsUnicode returns the unicode string for a TLD.
// Example: TLD xn--vermgensberater-ctb -> vermögensberater
func (t TLD) AsUnicode() string {
	i := idna.New(idna.MapForLookup(), idna.StrictDomainName(true))

	unicode, err := i.ToUnicode(t.String())
	if err != nil {
		return t.String()
	}

	return unicode
}

// AsPunycode returns the punycode sring for a TLD.
// Example: vermögensberater -> xn--vermgensberater-ctb
func (t TLD) AsPunycode() string {
	i := idna.New(idna.MapForLookup(), idna.StrictDomainName(true))

	punycode, err := i.ToASCII(t.String())
	if err != nil {
		return t.String()
	}

	return punycode
}

// IsValid will validate a TLD by passing a string.
func IsValid(s string) bool {
	return FromString(s).IsValid()
}

// FromDomainName returns the TLD from a host name or domain name.
func FromDomainName(name string) TLD {
	parts := strings.Split(name, ".")

	if len(parts) < 2 {
		return TLD("")
	}

	return TLD(parts[len(parts)-1])
}

// FromString returns the TLD from a string.
func FromString(s string) TLD {
	return TLD(s)
}
