package tld

import (
	"fmt"
	"testing"
)

func TestTLD_AsUnicode(t *testing.T) {
	cases := []struct {
		in  TLD
		out string
	}{
		{
			in:  TLD("se"),
			out: "se",
		},
		{
			in:  TLD("XN--VERMGENSBERATER-CTB"),
			out: "vermögensberater",
		},
		{
			in:  TLD("vermögensberater"),
			out: "vermögensberater",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			got := tc.in.AsUnicode()

			if got != tc.out {
				t.Fail()
			}
		})
	}
}
func TestTLD_AsPunycode(t *testing.T) {
	cases := []struct {
		in  TLD
		out string
	}{
		{
			in:  FromString("se"),
			out: "se",
		},
		{
			in:  FromString("XN--VERMGENSBERATER-CTB"),
			out: "xn--vermgensberater-ctb",
		},
		{
			in:  FromString("vermögensberater"),
			out: "xn--vermgensberater-ctb",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			got := tc.in.AsPunycode()

			if got != tc.out {
				t.Fail()
			}
		})
	}
}

func TestTLD_FromDomainName(t *testing.T) {
	cases := []struct {
		in  string
		out TLD
	}{
		{
			in:  "",
			out: FromString(""),
		},
		{
			in:  ".",
			out: FromString(""),
		},
		{
			in:  "github.com",
			out: FromString("com"),
		},
		{
			in:  "www.my-blog.xn--vermgensberater-ctb",
			out: FromString("xn--vermgensberater-ctb"),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			tld := FromDomainName(tc.in)

			if tld != tc.out {
				t.Fail()
			}
		})
	}
}

func TestTLD_Strings(t *testing.T) {
	cases := []struct {
		in   TLD
		outU string
		outL string
	}{
		{
			in:   FromString("com"),
			outU: "COM",
			outL: "com",
		},
		{
			in:   FromString("xn--vermgensberater-ctb"),
			outU: "XN--VERMGENSBERATER-CTB",
			outL: "xn--vermgensberater-ctb",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			if u := tc.in.String(); u != tc.outU {
				t.Fail()
			}

			if l := tc.in.LowerString(); l != tc.outL {
				t.Fail()
			}
		})
	}
}

func TestTLD_IsValid(t *testing.T) {
	cases := []struct {
		in    TLD
		valid bool
	}{
		{
			in:    FromString("apa"),
			valid: false,
		},
		{
			in:    FromString("APA"),
			valid: false,
		},
		{
			in:    FromString("se"),
			valid: true,
		},
		{
			in:    FromString("SE"),
			valid: true,
		},
		{
			in:    FromString("sE"),
			valid: true,
		},
		{
			in:    SE,
			valid: true,
		},
		{
			in:    FromString(SE.LowerString()),
			valid: true,
		},
		{
			in:    FromString("xn--apa-ctb"),
			valid: false,
		},
		{
			in:    FromString("xn--vermgensberater-ctb"),
			valid: true,
		},
		{
			in:    FromString("XN--VERMGENSBERATER-CTB"),
			valid: true,
		},
		{
			in:    FromString("XN--vermgensberater-CTB"),
			valid: true,
		},
		{
			in:    FromString("vermögensberater"),
			valid: true,
		},
		{
			in:    FromString("VERMÖGENSBERATER"),
			valid: true,
		},
		{
			in:    FromString("vermögensBERATER"),
			valid: true,
		},
		{
			in:    FromString("öermögensberateö"),
			valid: false,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			if tc.in.IsValid() != tc.valid {
				t.Fail()
			}
		})
	}
}

func TestTLD_IsValidString(t *testing.T) {
	cases := []struct {
		in    string
		valid bool
	}{
		{
			in:    "apa",
			valid: false,
		},
		{
			in:    "APA",
			valid: false,
		},
		{
			in:    "se",
			valid: true,
		},
		{
			in:    "SE",
			valid: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			if IsValid(tc.in) != tc.valid {
				t.Fail()
			}
		})
	}
}
