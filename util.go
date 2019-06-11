package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io"
	"bytes"
)

type Entry struct {
	Site  string
	Uname string
	Pw    string
}

func openOrCreate(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0644)
}

func loadFile(f *os.File) []Entry {
	dec := json.NewDecoder(f)
	lineCount, err := lineCounter(f) 
	if err != nil {
		fmt.Println("error counting lines")
		os.Exit(1)
	}
	fmt.Println(lineCount)
	entries := make([]Entry, lineCount*2)

	for dec.More() {
		var v Entry
		fmt.Println(v)
		if err := dec.Decode(&v); err != nil {
			fmt.Println("error parsing file")
			os.Exit(1)
		}
		entries = append(entries, v)
	}
	return entries
}

func writeToFile() {}

// works only once; maybe some sort of static state???
// one second invocation dec.More() returns false
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

// probaly same issue as entryExists
func getEntry(f *os.File, site string) *Entry {
	dec := json.NewDecoder(f)
	

	for dec.More() {
		var v Entry
		if err := dec.Decode(&v); err != nil {
			fmt.Println("error parsing database")
			os.Exit(1)
		}

		if v.Site == site {
			return &v
		}	
	}

	return nil
}

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count+1, nil

        case err != nil:
            return count+1, err
        }
    }
}