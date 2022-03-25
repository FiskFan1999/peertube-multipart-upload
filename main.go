package main

import (
	"log"
	"os"
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
