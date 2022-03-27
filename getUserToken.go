package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type OauthClientsLocal struct {
	Client_id     string
	Client_secret string
}

type PeertubeUserToken struct {
	Access_token string
}

func GetClientLocal(hostname string) (*OauthClientsLocal, error) {
	/*
		Get oauth local client via
		https://docs.joinpeertube.org/api-rest-getting-started?id=get-client
	*/
	clientslocalurl := fmt.Sprintf("%s/api/v1/oauth-clients/local", hostname)
	clientslocalreq, err := http.Get(clientslocalurl)
	defer clientslocalreq.Body.Close()
	if clientslocalreq.StatusCode != 200 {
		switch clientslocalreq.StatusCode {
		case 423:
			return nil, ErrorRateLimited
		default:
			return nil, errors.New(fmt.Sprintf("clientslocalreq url %s returned status code %d.", clientslocalurl, clientslocalreq.StatusCode))
		}
	}
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(clientslocalreq.Body)
	if err != nil {
		return nil, err
	}

	local := new(OauthClientsLocal)

	if err = json.Unmarshal(body, local); err != nil {
		return nil, err
	}

	return local, nil
}

func GetUserTokenFromAPI(hostname, username, password string) (string, error) {
	clientlocal, err := GetClientLocal(hostname)
	if err != nil {
		return "", err
	}

	postForm := url.Values{}
	postForm.Add("client_id", clientlocal.Client_id)
	postForm.Add("client_secret", clientlocal.Client_secret)
	postForm.Add("grant_type", "password")
	postForm.Add("response_type", "code")
	postForm.Add("username", username)
	postForm.Add("password", password)

	gutUrl := fmt.Sprintf("%s/api/v1/users/token", hostname)
	tokenreq, err := http.PostForm(gutUrl, postForm)
	defer tokenreq.Body.Close()

	if tokenreq.StatusCode != 200 {
		switch tokenreq.StatusCode {
		case 423:
			return "", ErrorRateLimited
		default:
			return "", errors.New(fmt.Sprintf("tokenreq URL %s returned status code %d.", gutUrl, tokenreq.StatusCode))
		}
	}

	if err != nil {
		return "", err
	}
	tokenreqbody, err := io.ReadAll(tokenreq.Body)
	if err != nil {
		return "", err
	}
	tokenStr := new(PeertubeUserToken)
	if err = json.Unmarshal(tokenreqbody, tokenStr); err != nil {
		return "", err
	}
	return tokenStr.Access_token, nil
}
