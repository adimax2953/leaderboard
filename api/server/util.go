package server

// package util
// package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"sync"
	"time"
)

// func main() {
// 	test, err := NewAESEncoder("111111111111111111111111")
// 	if err != nil {
// 		panic(err)
// 	}

// 	tsencrypt, err := test.Encrypt("456789")
// 	fmt.Println(tsencrypt, err)

// 	tsdecrypt, err := test.Decrypt(tsencrypt)
// 	fmt.Println(tsdecrypt, err)
// }

// AESEncoder - AES 編碼器
type AESEncoder struct {
	Black              cipher.Block
	sSQLClientSyncOnce sync.Once
	Key                []byte
}

// NewAESEncoder - The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func NewAESEncoder(key string) (*AESEncoder, error) {
	var err error
	aesEncoder := &AESEncoder{}
	if key == "" {
		err = errors.New("'key' cannot be empty")
		return aesEncoder, err
	}

	aesEncoder.sSQLClientSyncOnce.Do(func() {
		aesEncoder.Key = []byte(key)
		aesEncoder.Black, err = aes.NewCipher(aesEncoder.Key)
	})
	if err != nil {
		return aesEncoder, err
	}

	return aesEncoder, nil
}

// Encrypt -
func (aese *AESEncoder) Encrypt(plaintext string) (string, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cipher.NewCFBEncrypter(aese.Black, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		[]byte(plaintext))
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt -
func (aese *AESEncoder) Decrypt(d string) ([]byte, error) {
	ciphertext, err := hex.DecodeString(d)
	if err != nil {
		return []byte{}, err
	}
	if len(ciphertext) < aes.BlockSize {
		return []byte{}, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cipher.NewCFBDecrypter(aese.Black, iv).XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// getTime -
func getTime(timestamp int64) *time.Time {
	if timestamp >= 0 {
		t := time.Unix(timestamp, 0)
		return &t
	}
	return nil
}
