package main

import (
	"errors"
)

var (
	ErrorRateLimited = errors.New("Request returned status code 423 Too Many Requests")
)
