package main

import "os"

func GetDescriptionTextFromFilename(filename string) (string, error) {
	/*
		Wrapper of io function
	*/
	b, e := os.ReadFile(filename)
	return string(b), e
}
