package tld

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Validator is the type that knows how to get and store the list of all
// currently available TLDs from IANA.
type Validator struct {
	client  *http.Client
	offline bool
	mu      sync.RWMutex
	tlds    map[TLD]struct{}
}

// NewValidator will create a new validator with a default client and do an
// initial refresh to get all the current TLDs.
func NewValidator() *Validator {
	v := &Validator{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		mu:      sync.RWMutex{},
		offline: false,
		tlds:    map[TLD]struct{}{},
	}

	if err := v.Refresh(); err != nil {
		panic(err)
	}

	return v
}

// Refresh will perofrm a HTTP request to the IANA web page listing all
// currently approved and valid TLDs. Refresh may be called how many times you
// like and will always to a complete reset of the TLD list.
func (v *Validator) Refresh() error {
	response, err := v.client.Get(ianaURL)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("could not perform request")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	result := strings.Split(string(body), "\n")

	v.mu.Lock()
	defer v.mu.Unlock()

	// Reset TLD list to clear old entries.
	v.tlds = map[TLD]struct{}{}

	for _, tld := range result[1 : len(result)-1] {
		v.tlds[TLD(tld)] = struct{}{}
	}

	return nil
}

// IsValid returns true or false telling if a TLD is valid. The method is case
// insensitive and can handle strings, TLDs and fmt.Stringers.
func (v *Validator) IsValid(tld interface{}) bool {
	var t TLD

	switch v := tld.(type) {
	case string:
		t = TLD(v)
	case fmt.Stringer:
		t = TLD(v.String())
	default:
		return false
	}

	v.mu.RLock()
	defer v.mu.RUnlock()

	// Convert to uppercase punycode before matching.
	punycode := strings.ToUpper(t.AsPunycode())

	_, ok := v.tlds[TLD(punycode)]

	return ok
}
