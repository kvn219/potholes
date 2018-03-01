package main

import (
	"testing"
)

// TestGetRequest checks the 200 status of our get request made to Socrata.
func TestGetRequest(t *testing.T) {
	resp := getReqFromSocrata()
	if resp.StatusCode != 200 {
		t.Fail()
	}
}
