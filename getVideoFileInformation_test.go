package main

import (
	"errors"
	"os"
	"strings"
	"testing"
)

type GetFileResult struct {
	File *os.File
	Err  error
}

var (
	GetFileCases = map[string]GetFileResult{
		"<filename>":             {new(os.File), nil},
		"  <filename>":           {new(os.File), nil},
		"<filename> ":            {new(os.File), nil},
		"/tmp/someotherfile.mp4": {nil, os.ErrNotExist},
	}
)

func TestGetFileFromFilename(t *testing.T) {
	file, err := os.CreateTemp("", "video.*.mp4")
	defer file.Close()
	defer os.Remove(file.Name())
	if err != nil {
		t.Errorf("Error when setting up TestGetFileFromFilename: %s", err.Error())
	}
	fname := file.Name()
	t.Log(fname)

	for key, val := range GetFileCases {
		keycorrect := strings.ReplaceAll(key, "<filename>", fname)
		var result GetFileResult
		result.File, result.Err = GetFileFromFilename(keycorrect)

		if (result.File != nil && val.File == nil) || (result.File == nil && val.File != nil) {
			t.Errorf("For filename %s, should have returned %v but instead returned %+v.", keycorrect, val.File, result.File)
		}
		if !errors.Is(result.Err, val.Err) {
			t.Errorf("For filename %s, should have recieved error %s but instead recieved %+v", keycorrect, val.Err, result.Err)
		}
	}

}
