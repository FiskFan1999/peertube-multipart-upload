package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestHelpText(t *testing.T) {
	file, err := os.Open("help.txt")
	if err != nil {
		t.Errorf("Error while opening help.txt: \"%+v\"\n", err)
		return
	}

	filecont, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("Error while reading file: \"%+v\"\n", err)
		return
	}
	if !reflect.DeepEqual(filecont, HelpText) && !(len(filecont) == 0 && len(HelpText) == 0) {
		t.Error("HelpText should be the contents of help.txt.")
	}
}
