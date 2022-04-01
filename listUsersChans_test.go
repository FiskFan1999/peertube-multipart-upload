package main

import (
	"testing"
)

var FramesoftTest []UserChannel = []UserChannel{
	UserChannel{"libre.culture", 623},
	UserChannel{"mooc.chatons.1", 370},
	UserChannel{"e4985792-98ca-49be-a1aa-bceecd1c8051", 26},
	UserChannel{"bf54d359-cfad-4935-9d45-9d6be93f63e8", 25},
}

func TestGetChannelsForUser(t *testing.T) {
	var resp *APIAccountsChannelsResponse
	resp, err := GetChannelsForUser("framasoft@framatube.org", TESTSERVERHOSTNAME)
	if err != nil {
		t.Error(err.Error())
		return
	}

	/*
		Test that the returned JSON is correct.
	*/
	if len(resp.Data) != len(FramesoftTest) {
		t.Error("Length of api returned channel list != expected channel list")
	}

	for i := range resp.Data {
		a := resp.Data[i]
		b := FramesoftTest[i]
		if a.Name != b.Name {
			t.Errorf("Recieved channel name %s != Expected channel name %s for index %d", a.Name, b.Name, i)
		}
		if a.Id != b.Id {
			t.Errorf("Recieved channel Id %d != Expected channel Id %d for index %d", a.Id, b.Id, i)
		}
	}
}
