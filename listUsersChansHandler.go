package main

import (
	"log"
	"os"
)

func ListUserChansHandler() {
	hostnamer, hnok := os.LookupEnv("PTHOST")
	username, unok := os.LookupEnv("PTUSER")

	if !hnok || !unok {
		log.Println("Environment variables $PTHOST and $PTUSER required.")
		log.Printf("PTHOST='https://example.net' PTUSER=username %s channels", os.Args[0])
		return
	}
	hostname := CleanHostname(hostnamer)

	var channels *APIAccountsChannelsResponse
	channels, err := GetChannelsForUser(username, hostname)
	if err != nil {
		log.Printf("Error: %+v", err)
		return
	}

	log.Printf("%d total channels:", channels.Total)
	log.Println()
	log.Println("#    ID\t Name")
	var i int
	var ch UserChannel
	for i, ch = range channels.Data {
		log.Printf("%d   %d \t%s", i, ch.Id, ch.Name)
	}
	log.Println()

}
