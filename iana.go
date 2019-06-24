package tld

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	ianaURL = "http://data.iana.org/TLD/tlds-alpha-by-domain.txt"
)

// IANA is the type that knows how to get and store the list of all currently
// available TLDs from IANA.
type IANA struct {
	client  *http.Client
	offline bool
	mu      sync.RWMutex
	tlds    map[TLD]struct{}
}

// NewIANA will create a new IANA with a default client and do an initial
// refresh to get all the current TLDs.
func NewIANA() *IANA {
	v := &IANA{
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

// Refresh will perform a HTTP request to the IANA web page listing all
// currently approved and valid TLDs. Refresh may be called how many times you
// like and will always to a complete reset of the TLD list.
func (v *IANA) Refresh() error {
	response, err := v.client.Get(ianaURL)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Print(err.Error())
		}
	}()

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

// All returns all known TLDs according to IANA. The list returned will be empty
// if Refresh() has not been invoked.
func (v *IANA) All() []TLD {
	v.mu.RLock()
	defer v.mu.RUnlock()

	tlds := []TLD{}

	for tld := range v.tlds {
		tlds = append(tlds, tld)
	}

	sort.Slice(tlds, func(i, j int) bool {
		return tlds[i].String() < tlds[j].String()
	})

	return tlds
}

// IsValid returns true or false telling if a TLD is valid. The method is case
// insensitive and can handle strings, TLDs and fmt.Stringers.
func (v *IANA) IsValid(tld interface{}) bool {
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
