package main

import (
	"testing"
)

var (
	CleanHostnameCases = map[string]string{
		"":                     "",
		"https://example.net/": "https://example.net",
		"https://example.net":  "https://example.net",
	}
)

func TestCleanHostname(t *testing.T) {
	for key, val := range CleanHostnameCases {
		out := CleanHostname(key)
		if out != val {
			t.Errorf("For original hostname \"%+v\", expected result \"%+v\" but got \"%+v\".", key, val, out)
		}
	}
}
