package main

import (
	"log"
	"os"
	"strings"
)

var (
	AppData, _ = GetConfigDir()
)

func GetConfigDir() (string, error) {
	/*
		Wrapper around os.UserConfigDir which
		exits the entire program if the
		function returns an error for
		any reason.
	*/
	appData, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Attempt to find user config directory returned the following error: \"%s\"\n", err.Error())
	}
	return strings.TrimRight(appData, "/"), err
}
