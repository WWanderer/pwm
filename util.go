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

func newDatabase(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
}

func entryExists(f *os.File, e Entry) bool {
	dec := json.NewDecoder(f)

	for dec.More() {
		var v Entry
		if err := dec.Decode(&v); err != nil {
			fmt.Println("error parsing database")
			os.Exit(1)
		}

		if v.Site == e.Site && v.Uname == e.Uname {
			return true
		}
	}

	return false
}
