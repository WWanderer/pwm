package main

import (
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/scrypt"
)
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// possible switch to https://github.com/gtank/cryptopasta/blob/master/encrypt.go
// because these don't fail on wrong password

func Encrypt(json, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(json))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], json)
	return ciphertext
}

func Decrypt(ciphertext, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plaintext := make([]byte, len(ciphertext))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return plaintext
}

// https://godoc.org/github.com/sethvargo/go-password/password
func genPW(length int) string {
	password, err := password.Generate(length, length/4, length/4,
		false, true)
	if err != nil {
		panic(err)
	}
	return password
}

// https://godoc.org/golang.org/x/crypto/scrypt
func aesKey(pw []byte) []byte {
	salt := []byte("salt")
	hash, err := scrypt.Key(pw, salt, 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return hash
}
