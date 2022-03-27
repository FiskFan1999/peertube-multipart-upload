package main

import (
	"errors"
	"os"
	"testing"
)

func TestGetUserTokenFromAPI(t *testing.T) {
	hostnamer, hnok := os.LookupEnv("PTHOST")
	username, unok := os.LookupEnv("PTUSER")
	password, pwok := os.LookupEnv("PTPASSWD")
	if !hnok || !unok || !pwok {
		/*
			At least one of the above is not
			specified. Skip.
		*/
		t.Log("Skipping user token test due to no user login provided.")
		t.Skip()
	}
	hostname := CleanHostname(hostnamer)
	token, err := GetUserTokenFromAPI(hostname, username, password)
	if err != nil {
		if errors.Is(err, ErrorRateLimited) {
			t.Log("User token generation returned 423, is rate limited. Skipping.")
			t.Skip()
		}
		t.Errorf("GetUserTokenFromAPI error %+v", err)
	}
	t.Logf("Token recieved, len=%d", len(token))

	_, err = GetUserTokenFromAPI(hostname, username, password+"wrong")
	if err == nil {
		t.Error("On oauth request with incorrect password, GetUserTokenFromAPI returned true.")
	} else {
		if errors.Is(err, ErrorRateLimited) {
			t.Log("User token generation returned 423, is rate limited. Skipping.")
			t.Skip()
		} else if errors.Is(err, IncorrectOauthLogin) {
			t.Log("Oauth request correctly returned status code 400 bad request on incorrect password")
		} else {
			t.Errorf("GetUserTokenFromAPI error %+v (should be 400)", err)
		}
	}
}
