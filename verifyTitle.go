package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TitleMinLength = 3
	TitleMaxLength = 120
)

var (
	TitleLengthError = errors.New(fmt.Sprintf("Title shorter than %d or longer than %d characters.", TitleMinLength, TitleMaxLength))
)

func VerifyTitle(input string) (string, error) {
	strip := strings.TrimSpace(input)

	if l := len(strip); l < TitleMinLength || l > TitleMaxLength {
		return "", TitleLengthError
	}

	return strip, nil
}
