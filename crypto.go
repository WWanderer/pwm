package main

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
	"os"
)
// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
func Encrypt(entries []Entry) {}

func Decrypt(entries []Entry) {}

func genPW(length int) (string, error) {
	// Generate a password that is length characters long with
	// length/4 digits, length/4 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	password, err := password.Generate(length, length/4, length/4, false, true)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return password, nil
}
