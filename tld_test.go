package tld

import (
	"fmt"
	"testing"
)

func TestValidator_AsUnicode(t *testing.T) {
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
func TestValidator_AsPunycode(t *testing.T) {
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

func TestValidator_FromDomainName(t *testing.T) {
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
