package main

import (
	"reflect"
	"strings"
	"testing"
)

type VTResult struct {
	Output string
	Err    error
}

var VTCases = map[string]VTResult{
	"hello":             {"hello", nil},
	" this is a title":  {"this is a title", nil},
	"this is a title  ": {"this is a title", nil},

	"":   {"", TitleLengthError},
	"sh": {"", TitleLengthError},

	strings.Repeat("a", 121):        {"", TitleLengthError},
	strings.Repeat("a", 120):        {strings.Repeat("a", 120), nil},
	"  " + strings.Repeat("a", 120): {strings.Repeat("a", 120), nil},
}

func TestVerifyTitle(t *testing.T) {
	for key, val := range VTCases {
		var result VTResult
		result.Output, result.Err = VerifyTitle(key)
		if !reflect.DeepEqual(result, val) {
			t.Errorf("For input %s VerifyTitle expected %v but instead recieved %v.", key, val, result)
		}
	}
}
