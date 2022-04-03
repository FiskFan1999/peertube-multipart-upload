package main

import (
	"log"
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
		case "help":
			FullHelpHandler()
		default:
			log.Printf("unknown arg %s\n", os.Args[1])
		}
		return
	}

	/*
		Run the multipart upload
	*/
	input, err, failtext := ReadEnvironmentVars()
	if err != nil {
		log.Println(strings.Join(failtext, "\n"))
		os.Exit(1)
	}
	log.Println()
	log.Println(input)

}
