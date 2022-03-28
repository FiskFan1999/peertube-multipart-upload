package main

import (
	"os"
	"strings"
)

const (
	TESTSERVERHOSTNAME = "https://peertube.cpy.re"
)

func main() {
	if len(os.Args) == 2 {
		//if strings.ToLower(os.Args[1]) == "list" {
		switch strings.ToLower(os.Args[1]) {
		case "list":
			ListUserChansHandler()
			return
		case "help":
			FullHelpHandler()
			return
		}
	}
}
