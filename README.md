# Go TLD Validator

A proper way to work with top level domains by validating towards
[IANA](https://data.iana.org/TLD/tlds-alpha-by-domain.txt) TLD list. Whenever a
new instance of a TLD validator is created the list will be fetched and may
after that be updated when desired. The validator is thread safe and can access
the valid storage concurrently.

## TLD Type

The TLD type makes it easy to convert between unicode and punycode for easier
handling and validation. A TLD may be retreived from a domain or host name by
using `FromDomanName()`.

## Usage

A generated file with all current TLDs from IANA can be used for offlien
validation. To update the list run `go generate ./...`.

```go
tld := FromDomanName("www.github.कॉम")

// Offline validation.
if !tld.IsValid() {
    fmt.Printf("Invalid TLD: %s\n", tld.AsPunycode())
}

// Online validation.
iana := NewIANA()

if !iana.IsValid(tld) {}
    fmt.Printf("Invalid TLD: %s\n", tld.AsPunycode())
}
```
