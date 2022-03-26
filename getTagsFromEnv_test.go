package main

import (
	"errors"
	"reflect"
	"testing"
)

var (
	TagsParseTestSuc = map[string][]string{
		"":              []string{},
		"one":           []string{"one"},
		"one,two":       []string{"one", "two"},
		"one,one":       []string{"one"},
		"one,one ":      []string{"one"},
		"one  ,two":     []string{"one", "two"},
		"one  , two ":   []string{"one", "two"},
		",one,two":      []string{"one", "two"},
		"one,two,":      []string{"one", "two"},
		"one,,two":      []string{"one", "two"},
		"one,two,three": []string{"one", "two", "three"},

		"one,two,three,four":      []string{"one", "two", "three", "four"},
		"one,two,three,four,five": []string{"one", "two", "three", "four", "five"},

		"one, two three, three, four, five": []string{"one", "two three", "three", "four", "five"},

		"123456789012345678901234567890": []string{"123456789012345678901234567890"},
	}

	TagsParseTestTooMany = []string{
		"taga,tagb,tagc,tagd,tage,tagf",
		"tagg,taga,tagb,tagc,tagd,tage,tagf",
		"taggh,taga,tagb,tagc,tagd,tage,tagf",
		"taga,tagb,tagc,tagd,tage,taga,tagb,tagc,tagd,tage",
	}

	TagsParseTestTooLong = map[string]int{
		"1234567890123456789012345678901":                       0,
		"02345,1234567890123456789012345678901":                 1,
		"02345,23451,1234567890123456789012345678901,12334":     3,
		"0123,1243,22345,1234567890123456789012345678901,12345": 3,

		"02345,1234567890123456789012345678901,123,1234567890123456789012345678901": 1,

		"1,tag1,tag3,tag4,tag5": 0,
		"tag1,tag2,t,tag4":      2,
		"tag1,tag2,t,t":         2,
	}
)

func TestGetTagsFromEnv(t *testing.T) {
	/*
		Test cases which should not return
		an error.
	*/
	for key, val := range TagsParseTestSuc {
		testRes, err, i := GetTagsFromEnv(key)
		if err != nil {
			t.Errorf("For input %s, should not have returned error, but got error \"%s\", error index value %v.", key, err.Error(), i)
		}
		if i != nil {
			t.Errorf("for input %s, error index should be nil, but got %d.", key, i)
		}

		if !reflect.DeepEqual(testRes, val) && !(len(testRes) == 0 && len(val) == 0) {
			// Note, reflect.DeepEqual = false if comparing two empty slices
			// result is not equal
			t.Errorf("For input %s, expected result %s, got value %s", key, val, testRes)
		}

	}

	/*
		Test cases which should return an error
	*/
	for _, val := range TagsParseTestTooMany {
		testRes, err, i := GetTagsFromEnv(val)
		if !errors.Is(err, TooManyTagsError) {
			t.Errorf("For input %s, should have recieved TooManyTagsError, but recieved \"%s\" instead.", val, err)
		}

		if len(testRes) != 0 {
			t.Errorf("For input %s, returned slice should have been empty, but recieved %v.", val, i)

		}

		if i != nil {
			t.Errorf("For input %s, error index should be nil, but is %d.", val, i)
		}
	}

	for key, val := range TagsParseTestTooLong {
		testRes, err, i := GetTagsFromEnv(key)
		if !errors.Is(err, TagTooLongError) {
			t.Errorf("For input %s, error should be TagTooLongError but is actually %s.", key, err)
		}
		if i == nil {
			t.Errorf("For input %s, error index should be %d but is actually nil.", key, val)
		} else if val != *i {
			t.Errorf("For input %s, error index should be %d but is actually %v.", key, val, *i)
		}

		if len(testRes) != 0 {
			t.Errorf("For input %s, tags slice should be empty but is actually %v.", key, testRes)
		}
	}
}
