package main

import (
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/scrypt"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"io"
	// "os"
)

// http://www.golangprograms.com/cryptography/advanced-encryption-standard.html
func Encrypt(encodedText []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	ciphertext := make([]byte, aes.BlockSize+len(encodedText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], encodedText)
	return ciphertext
}

func Decrypt(entries []Entry) {}

// https://godoc.org/github.com/sethvargo/go-password/password
func genPW(length int) (string, error) {
	password, err := password.Generate(length, length/4, length/4,
		false, true)
	return password, err
}

// https://godoc.org/golang.org/x/crypto/scrypt
func aesKey(pw []byte) ([]byte, error) {
	salt := []byte("salt")
	hash, err := scrypt.Key(pw, salt, 32768, 8, 1, 32)
	return hash, err
}
