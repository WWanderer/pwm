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
	"os"
	"io/ioutil"
)

func Encrypt(json []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(json))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], json)
	return ciphertext
}

// https://gist.github.com/stupidbodo/601b68bfef3449d1b8d9
func Decrypt(f *os.File, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	iv := ciphertext[:aes.BlockSize]
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	return plaintext
}

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
