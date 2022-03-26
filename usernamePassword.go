package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/scrypt"
)

const (
	KeyLength = 32
	SaltLen   = 32
)

var (
	DecryptFailed = errors.New("cipher: message authentication failed")
)

func CreateKeyFromPassword(passwd string, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, SaltLen)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	result, err := scrypt.Key([]byte(passwd), salt, 16384, 8, 1, KeyLength)
	return result, salt, err
}

func EncryptStr(data, passwd string) ([]byte, error) {
	key, salt, err := CreateKeyFromPassword(passwd, nil)
	if err != nil {
		return nil, err
	}

	blockcipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockcipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func DecryptStr(data []byte, passwd string) (string, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, _, err := CreateKeyFromPassword(passwd, salt)
	if err != nil {
		return "", err
	}

	blockcipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(blockcipher)
	if err != nil {
		return "", err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil

}
