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
			in:  TLD("se"),
			out: "se",
		},
		{
			in:  TLD("XN--VERMGENSBERATER-CTB"),
			out: "xn--vermgensberater-ctb",
		},
		{
			in:  TLD("vermögensberater"),
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
			out: TLD(""),
		},
		{
			in:  ".",
			out: TLD(""),
		},
		{
			in:  "github.com",
			out: TLD("com"),
		},
		{
			in:  "www.my-blog.xn--vermgensberater-ctb",
			out: TLD("xn--vermgensberater-ctb"),
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
			in:   TLD("com"),
			outU: "COM",
			outL: "com",
		},
		{
			in:   TLD("xn--vermgensberater-ctb"),
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
			in:    TLD("apa"),
			valid: false,
		},
		{
			in:    TLD("APA"),
			valid: false,
		},
		{
			in:    TLD("se"),
			valid: true,
		},
		{
			in:    TLD("SE"),
			valid: true,
		},
		{
			in:    TLD("sE"),
			valid: true,
		},
		{
			in:    SE,
			valid: true,
		},
		{
			in:    TLD(SE.LowerString()),
			valid: true,
		},
		{
			in:    TLD("xn--apa-ctb"),
			valid: false,
		},
		{
			in:    TLD("xn--vermgensberater-ctb"),
			valid: true,
		},
		{
			in:    TLD("XN--VERMGENSBERATER-CTB"),
			valid: true,
		},
		{
			in:    TLD("XN--vermgensberater-CTB"),
			valid: true,
		},
		{
			in:    TLD("vermögensberater"),
			valid: true,
		},
		{
			in:    TLD("VERMÖGENSBERATER"),
			valid: true,
		},
		{
			in:    TLD("vermögensBERATER"),
			valid: true,
		},
		{
			in:    TLD("öermögensberateö"),
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
