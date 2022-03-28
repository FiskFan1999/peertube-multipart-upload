package main

import (
	"fmt"
)

var HelpText []byte = []byte("hello.")

func FullHelpHandler() {
	fmt.Printf("\n%s\n", HelpText)
}
