package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateEntry(f *os.File) {
	var tmp Entry

	fmt.Printf("site name\n~> ")
	_, err := fmt.Scan(&tmp.Site)
	if err != nil {
		fmt.Println("Error writing site name")
		return
	}

	fmt.Printf("username\n~> ")
	_, err = fmt.Scan(&tmp.Uname)
	if err != nil {
		fmt.Println("Error writing username")
		return
	}

	// TODO suggest either user entry or pw generation
	fmt.Printf("password\n~> ")
	_, err = fmt.Scan(&tmp.Pw)
	if err != nil {
		fmt.Println("Error writing password")
		return
	}

	exists := entryExists(f, tmp)
	if exists {
		fmt.Println("Entry already exists")
		return
	}

	buff, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println("Error encoding to json")
		return
	}
	buff = append(buff, '\n')

	// assumes flag os.O_APPEND is set on the file
	_, err = f.Write(buff)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}

}

func ReadEntry(f *os.File, site string) {
	e := getEntry(f, site)
	if e == nil {
		fmt.Println("Entry not found")
		return
	}

	fmt.Printf("Site: %s\nUsername: %s\nPassword: %s\n", e.Site, e.Uname, e.Pw)
}

func UpdateEntry(f *os.File, site string) {
	e := getEntry(f, site)
	if e == nil {
		fmt.Println("Entry not found")
		return
	}
	// https://golang.org/pkg/io/#example_SectionReader_Seek
	// calculate start position of entry
	// might be easier to use delete then create
}

func DeleteEntry(f *os.File, site string) {

}
