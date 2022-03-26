package main

import (
	"testing"
)

func TestCreateKeyFromPassword(t *testing.T) {
	for i := 0; i < 5; i++ {
		result, salt, err := CreateKeyFromPassword("mypass", nil)
		if err != nil {
			t.Errorf("CreateKeyFromPassword result %+v salt %+v error %+v", result, salt, err)
		} else {
			t.Logf("CreateKeyFromPassword key %+v", result)
		}
	}
}

func BenchmarkCreateKeyFromPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateKeyFromPassword("benchmarkpass", nil)
	}
}

func TestEncryptionDecryption(t *testing.T) {
	password := "letmeseemystuff"
	data := "Hello. This is the data."
	encrypted, err := EncryptStr(data, password)
	if err != nil {
		t.Errorf("Encryption function returned the following error: %+v", err)
	}

	t.Logf("encrypted text of %s: %+v", data, encrypted)

	decrypted, err := DecryptStr(encrypted, password)
	if err != nil {
		t.Errorf("Decryption function returned the following error: %s", err)
	}

	if decrypted != data {
		t.Error("Encryption and decryption did not succeed.")
	}

	// test if the wrong password is entered
	encrypted, err = EncryptStr(data, password)
	if err != nil {
		t.Errorf("Encryption function returned the following error: %+v", err)
	}

	decrypted, err = DecryptStr(encrypted, password+"wrong")
	if err == nil {
		t.Error("On decryption with wrong password, should have passed error but didn't.")
	} else if err != nil && err.Error() == DecryptFailed.Error() {
		t.Log("On failed password on decrypt, correctly passwed decryptfailederror")
	} else if decrypted != data {
		t.Error("Encryption and decryption did not succeed.")
	}
}

func BenchmarkEncryptionDecryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		en, _ := EncryptStr("This is the data.", "password")
		DecryptStr(en, "password")
	}

}
