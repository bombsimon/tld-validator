package tld

import (
	"strings"

	"golang.org/x/net/idna"
)

const (
	ianaURL = "http://data.iana.org/TLD/tlds-alpha-by-domain.txt"
)

// TLD represents a top level domain
type TLD string

// String returns the string value for a TLD.
func (t TLD) String() string {
	return string(t)
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

// FromDomainName returns the TLD from a host name or domain name.
func FromDomainName(name string) TLD {
	parts := strings.Split(name, ".")

	if len(parts) < 2 {
		return TLD("")
	}

	return TLD(parts[len(parts)-1])
}
