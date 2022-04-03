package main

import (
	"os"
	"strings"
)

func GetFileFromFilename(filename string) (file *os.File, err error) {
	trfilename := strings.TrimSpace(filename)
	file, err = os.Open(trfilename)
	return
}
