package main

import (
	"os"
	"strings"
)

const (
	TESTSERVERHOSTNAME = "https://peertube.cpy.re"
)

func main() {
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "list" {
		ListUserChansHandler()
		return
	}
}
