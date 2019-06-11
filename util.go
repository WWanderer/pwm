package main

import (
	"bufio"
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
