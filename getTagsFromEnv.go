package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TagMaxLength = 30
	TagMinLength = 2
	MaxTagsLimit = 5
)

var (
	TagTooLongError  = errors.New(fmt.Sprintf("Tag is greater than %d characters", TagMaxLength))
	TooManyTagsError = errors.New(fmt.Sprintf("More than %d tags specified", MaxTagsLimit))
)

func GetTagsFromEnv(s string) ([]string, error, *int) {
	/*
		This function also validates the
		string tags. For example, it checks
		that all of the tags are less than 30
		characters long.
	*/
	tags_raw := strings.Split(s, ",")

	/*
		strings.Strip tags and remove
		empty tags to sanitize input.
	*/
	var tags []string

	for _, rawt := range tags_raw {
		s := strings.TrimSpace(rawt)
		if len(s) == 0 {
			continue
		}
		shouldContinue := false
		for _, ss := range tags {
			if s == ss {
				// Dont enter the same tag twice if
				// entered by the user
				shouldContinue = true
				break
			}
		}
		if shouldContinue {
			continue
		}
		tags = append(tags, s)
	}

	// Note: check raw tags for length instead
	// of sanitized (shortened)
	if len(tags_raw) > MaxTagsLimit {
		// illegal
		return []string{}, TooManyTagsError, nil
	}

	for i, t := range tags {
		if len(t) < TagMinLength || len(t) > TagMaxLength {
			// illegal
			return []string{}, TagTooLongError, &i
		}
	}

	return tags, nil, nil
}
