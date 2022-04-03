package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGetDescriptionTextFromFilename(t *testing.T) {
	var tmpFileText string = "Hello. This is testing wether or not this text appears when the GetDescriptionTextFromFilename function is used."
	f, err := os.CreateTemp("", "*")
	if err != nil {
		t.Fatalf("Error while using os.CreateTemp: \"%+v\"\n", err)
	}
	defer os.Remove(f.Name())
	if _, err = f.WriteString(tmpFileText); err != nil {
		t.Fatalf("Error while writing to tmp file: %+v\n", err)
	}
	f.Close()

	txt, err := GetDescriptionTextFromFilename(f.Name())
	if err != nil {
		t.Fatalf("Error while reading from file via GetDescriptionTextFromFilename: %+v\n", err)
	}
	if !reflect.DeepEqual(txt, tmpFileText) {
		t.Fatal("Error: text returned by GetDescriptionTextFromFilename does not match that which was written to it.")
	}
}
