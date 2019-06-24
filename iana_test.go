package tld

import (
	"fmt"
	"testing"
)

func TestIANA_IsValid(t *testing.T) {
	cases := []struct {
		in    interface{}
		valid bool
	}{
		{
			in:    1,
			valid: false,
		},
		{
			in:    "",
			valid: false,
		},
		{
			in:    "unittest",
			valid: false,
		},
		{
			in:    "se",
			valid: true,
		},
		{
			in:    TLD("se"),
			valid: true,
		},
		{
			in:    "XN--VERMGENSBERATER-CTB",
			valid: true,
		},
		{
			in:    "xn--vermgensberater-ctb",
			valid: true,
		},
		{
			in:    "vermögensberater",
			valid: true,
		},
		{
			in:    "VERMÖGENSBERATER",
			valid: true,
		},
	}

	iana := NewIANA()

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			valid := iana.IsValid(tc.in)

			if valid != tc.valid {
				t.Fail()
			}
		})
	}
}
