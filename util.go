package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

type Entry struct {
	Site  string
	Uname string
	Pw    string
}

// json.Unmarshal could be useful for the rewrite
func loadFile(fileName string, key []byte) []Entry {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	var entries []Entry
	if !isEmpty(file) {
		content := Decrypt(file, key)

		reader := bytes.NewReader(content)
		dec := json.NewDecoder(reader)	

		for dec.More() {
			var e Entry
			if err := dec.Decode(&e); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			entries = append(entries, e)
		}
	}
	return entries
}

func writeFile(fileName string, entries []Entry, key []byte) error {
	var buffer [][]byte

	var file, err = os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	// TODO write separate function for json encoding
	for _, entry := range entries {
		e, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		e = append(e, '\n')
		buffer = append(buffer, e)
	}
	// TODO encrypt buffer
	for i := range buffer {
		encrypted := Encrypt(buffer[i], key)
		file.Write(encrypted)
	}
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

// return an empty entry in case of problem
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

	// TODO suggest either user entry or pw generation
	fmt.Printf("password\n~> ")
	scanner.Scan()
	tmp.Pw = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return Entry{}
	}

	return tmp
}

func isNil(e Entry) bool {
	if e.Site == "" && e.Uname == "" && e.Pw == "" {
		return true
	}
	return false
}

func isEmpty(f *os.File) bool {
	content, _ := ioutil.ReadAll(f)
	if len(content) != 0 {
		return false
	}
	return true
}
