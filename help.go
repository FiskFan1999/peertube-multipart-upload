package main

import (
	_ "embed"
	"fmt"
)

//go:embed help.txt
var HelpText []byte

func FullHelpHandler() {
	fmt.Printf("\n%s\n", HelpText)
}
