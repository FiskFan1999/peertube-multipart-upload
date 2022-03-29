package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGetFileSize(t *testing.T) {
	nano := time.Now().UnixNano()
	source := rand.NewSource(nano)
	gen := rand.New(source)
	for i := 0; i < 10; i++ {
		length := gen.Intn(1024 * 512)
		str := make([]byte, length)
		if _, err := gen.Read(str); err != nil {
			t.Errorf("Error while generating random file: \"%+v\"\n", err)
			return
		}
		file, err := os.CreateTemp("", "*.txt")
		if err != nil {
			t.Errorf("Error while generating random file: \"%+v\"\n", err)
			return
		}

		defer os.Remove(file.Name())

		_, err = file.Write(str)
		file.Close()
		if err != nil {
			t.Errorf("Error while writing to file: \"%+v\"\n", err)
			return
		}

		recievedLength, err := GetFileSize(file.Name())
		if err != nil {
			t.Errorf("Error during GetFileSize: \"%+v\"\n", err)
		}
		if recievedLength != int64(length) {
			t.Error("Error: GetFileSize did not return the correct length.")
		}

	}
}
