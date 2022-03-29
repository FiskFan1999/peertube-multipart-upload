package main

import (
	"os"
)

func GetFileSize(filename string) (int64, error) {
	file, err := os.Stat(filename)
	if err != nil {
		return int64(-1), err
	}
	return file.Size(), nil // size in bytes
}
