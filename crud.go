package main

import (
	"fmt"
)

func createEntry(fileName string, entries []Entry, key []byte) []Entry {
	tmp := buildEntry()
	if isNil(tmp) {
		fmt.Println("error reading your input")
		return entries
	}

	exists := entryExists(tmp, entries)
	if exists {
		fmt.Println("entry already exists")
		return entries
	}
	entries = append(entries, tmp)

	err := writeFile(fileName, entries, key)
	if err != nil {
		panic(err)
	}
	return entries
}

func readEntry(entries []Entry, site string) {
	for i := range entries {
		if entries[i].Site == site {
			fmt.Printf("Site: %s\nUsername: %s\nPassword: %s\n",
				entries[i].Site, entries[i].Uname, entries[i].Pw)
			return
		}
	}
	fmt.Println("entry not found")
}

func updateEntry(fileName string, entries []Entry, site string, key []byte) []Entry {
	updated := false
	exists := false

	for i := range entries {
		if entries[i].Site == site {
			exists = true
			e := entries[i]
			fmt.Println("input new information for `", e.Site, "`, leave blank to keep the same")
			var tmp Entry = buildEntry()

			if isNil(tmp) {
				fmt.Println("nothing changed")
				return entries
			}

			if tmp.Site != e.Site && tmp.Site != "" {
				entries[i].Site = tmp.Site
				updated = true
			}
			if tmp.Uname != e.Uname && tmp.Uname != "" {
				entries[i].Uname = tmp.Uname
				updated = true
			}
			if tmp.Pw != e.Pw && tmp.Pw != "" {
				entries[i].Pw = tmp.Pw
				updated = true
			}
			break
		}
	}

	if updated {
		err := writeFile(fileName, entries, key)
		if err != nil {
			fmt.Println(err)
		}
	}
	if !exists {
		fmt.Println("entry not found")
	}
	return entries
}

func deleteEntry(fileName string, entries []Entry, site string, key []byte) []Entry {
loop:
	for i := range entries {
		if entries[i].Site == site {
			fmt.Printf("are you sure you want to delete entry %s?\n (y/N)\n~> ", site)
			var confirm string
			fmt.Scan(&confirm)
			switch confirm {
			case "y":
				entries = append(entries[:i], entries[i+1:]...)
				fmt.Println("deleted entry ", site)
				err := writeFile(fileName, entries, key)
				if err != nil {
					panic(err)
				}
				break loop
			default:
				fmt.Println("did not delete entry for ", site)
				break loop
			}
		}
	}
	return entries
}

func printSites(entries []Entry) {
	fmt.Printf("Sites: \n")
	for i := range entries {
		if i%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s\t", entries[i].Site)
	}
	fmt.Println()
}
