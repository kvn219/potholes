package main

import (
	"testing"
)

// TestGetRequest checks the 200 status of our get request made to Socrata.
func TestGetRequest(t *testing.T) {
	url := "https://data.cityofnewyork.us/resource/fhrw-4uyv.json?$where=descriptor%20=%20%27Pothole%27&$limit=3&$order=created_date%20DESC"
	resp := getReqFromSocrata(url)
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestRequestUnreachable(t *testing.T) {
	url := "https://data.cityofnewyork.us/resource/unreachable!"
	resp := getReqFromSocrata(url)
	if resp.StatusCode != 404 {
		t.Fail()
	}
}
