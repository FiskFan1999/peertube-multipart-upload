package main

import (
	"errors"
	"os"
	"testing"
)

func TestGetConfigDir(t *testing.T) {
	/*
		Note: GetConfigDir panics on error
		from os.UserConfigDir() which is
		still considered a failed test.
	*/
	a, erra := GetConfigDir()
	b, errb := os.UserConfigDir()
	if !errors.Is(erra, errb) {
		t.Errorf("GetConfigDir returned incorrect error. Recieved \"%+v\", wanted \"%+v\".\n", erra, errb)
	} else if a != b {
		t.Errorf("GetConfigDir returned wrong directory. Recieved %s, wanted %s.\n", a, b)
	}

}
