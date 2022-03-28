package main

import (
	"reflect"
	"testing"
)

var GetVersionCases map[string]VersStrType = map[string]VersStrType{
	"{\"version\":\"vers\", \"commit\": \"comm\"}":  {"vers", "comm"},
	"{\"version\":\"vers2\", \"commit\": \"comm\"}": {"vers", "comm"},
}

func TestGetVersion(t *testing.T) {
	for key, val := range GetVersionCases {
		res, err := GetVersStr([]byte(key))
		if err != nil {
			t.Errorf("GetVersStr failed with following error: \"%+v\"\n", err)
		} else if !reflect.DeepEqual(*res, val) {
			t.Errorf("GetVersStr returned incorrect Version and Commit struct. Expected \"%+v\", got \"%+v\"\n", val, *res)
		}
	}
}
