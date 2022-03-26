package main

import (
	"log"
	"os"
)

const (
	TESTSERVERHOSTNAME = "https://peertube.cpy.re"
)

func main() {
	log.Println("peertube-multipart-upload")

	/*
		reading environment variables example:
	*/
	username, uok := os.LookupEnv("PTUNAME")
	log.Println("environment variable:")
	log.Println("PTUNAME, ok")
	log.Println(username, uok)

}
