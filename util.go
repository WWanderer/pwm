package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Entry struct {
	Site  string
	Uname string
	Pw    string
}

func loadFile(fileName string, key []byte) []Entry {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var entries []Entry

	ciphertext, empty := isEmpty(file)
	if !empty {
		content := Decrypt(ciphertext, key)

		reader := bytes.NewReader(content)
		dec := json.NewDecoder(reader)

		for dec.More() {
			var e Entry
			if err := dec.Decode(&e); err != nil {
				break
			}
			entries = append(entries, e)
		}
	}
	return entries
}

func writeFile(fileName string, entries []Entry, key []byte) error {
	var buffer bytes.Buffer

	var file, err = os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// TODO write separate function for json encoding
	for _, entry := range entries {
		e, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		_, err = buffer.Write(e)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	jsonBuf := buffer.Bytes()
	encrypted := Encrypt(jsonBuf, key)
	_, err = file.Write(encrypted)
	if err != nil {
		panic(err)
	}
	file.Sync()
	return nil
}

func entryExists(e Entry, entries []Entry) bool {
	for _, entry := range entries {
		if entry.Site == e.Site {
			return true
		}
	}
	return false
}

func buildEntry() Entry {
	var tmp Entry
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("site name\n~> ")
	scanner.Scan()
	tmp.Site = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return Entry{}
	}

	fmt.Printf("username\n~> ")
	scanner.Scan()
	tmp.Uname = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return Entry{}
	}

	fmt.Println("do you want to automatically generate a password? [y]/n")
	scanner.Scan()
	switch scanner.Text() {
	case "n":
		fmt.Printf("password\n~> ")
		scanner.Scan()
		tmp.Pw = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return Entry{}
		}
	default:
		tmp.Pw = genPW(16)
	}

	return tmp
}

func isNil(e Entry) bool {
	if e.Site == "" && e.Uname == "" && e.Pw == "" {
		return true
	}
	return false
}

func isEmpty(f *os.File) ([]byte, bool) {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	if len(content) != 0 {
		return content, false
	}
	return []byte(""), true
}
