package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIAccountsChannelsResponse struct {
	Total int
	Data  []UserChannel
}

type UserChannel struct {
	Name string
	Id   int
}

func GetChannelsForUser(username string, hostname string) (*APIAccountsChannelsResponse, error) {
	// https://peertube.cpy.re/api/v1/accounts/{name}/video-channels
	url := fmt.Sprintf("%s/api/v1/accounts/%s/video-channels", hostname, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *APIAccountsChannelsResponse = new(APIAccountsChannelsResponse)
	if err = json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return result, nil
}
