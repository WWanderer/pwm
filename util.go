package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	Site  string
	Uname string
	Pw    string
}

func openOrCreate(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
}

func loadFile(fileName string) []Entry {
	file, err := openOrCreate(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	var entries []Entry

	for dec.More() {
		var e Entry
		if err := dec.Decode(&e); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		entries = append(entries, e)
	}
	return entries
}

func writeToFile(fileName string, entries []Entry) error {
	var buffer [][]byte

	var file, err = os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	for _, entry := range entries {
		// encode each entry to json
		e, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		e = append(e, '\n')
		// put them in a buffer
		buffer = append(buffer, e)
	}
	// TODO encrypt buffer
	// write buffer to file
	for i := range buffer {
		file.Write(buffer[i])
	}
	return nil
}

func entryExists(e Entry, entries []Entry) bool {
	return false
}

// return an empty entry in case of problem
// TODO fix not accepting empty input
func buildEntry() Entry {
	var tmp Entry

	fmt.Printf("site name\n~> ")
	_, err := fmt.Scan(&tmp.Site)
	if err != nil {
		fmt.Println(err.Error())
		return Entry{}
	}

	fmt.Printf("username\n~> ")
	_, err = fmt.Scan(&tmp.Uname)
	if err != nil {
		fmt.Println(err.Error())
		return Entry{}
	}

	// TODO suggest either user entry or pw generation
	fmt.Printf("password\n~> ")
	_, err = fmt.Scan(&tmp.Pw)
	if err != nil {
		fmt.Println(err.Error())
		return Entry{}
	}
	return tmp
}

func isValid(e Entry) bool {
	if e.Site == "" && e.Uname == "" && e.Pw == "" {
		return false
	}
	return true
}
