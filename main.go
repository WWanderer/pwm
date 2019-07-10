package main

import "github.com/howeyc/gopass"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	fileName string
	key      []byte
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	entries := []Entry{}
	if len(os.Args) != 2 {
		var password []byte
		fileName, password = newFile()
		key = aesKey(password)
	} else {
		fileName = os.Args[1]
		fmt.Printf("enter your password\n~> ")
		pw, _ := gopass.GetPasswdMasked()
		key = aesKey(pw)
		entries = loadFile(fileName, key)
	}

	for {
		fmt.Printf("Enter a command:\n")
		fmt.Printf(
			"[c]reate\t[r]ead <entry>\t[u]pdate <entry>\t[d]elete <entry>\t[l]ist all sites\t[q]uit\n")
		scanner.Scan()
		command := []string{""}
		if scanner.Text() != "" {
			command = strings.Fields(scanner.Text())
		}
		switch command[0] {
		case "c":
			if len(command) != 1 {
				fmt.Println("usage: c")
				continue
			}
			entries = createEntry(fileName, entries, key)
		case "r":
			if len(command) != 2 {
				fmt.Println("usage: r <sitename>")
				continue
			}
			readEntry(entries, command[1])
		case "u":
			if len(command) != 2 {
				fmt.Println("usage: u <sitename>")
				continue
			}
			entries = updateEntry(fileName, entries, command[1], key)
		case "d":
			if len(command) != 2 {
				fmt.Println("usage: d <sitename>")
				continue
			}
			entries = deleteEntry(fileName, entries, command[1], key)
		case "l":
			if len(command) != 1 {
				fmt.Println("usage: l")
				continue
			}
			printSites(entries)
		case "q":
			os.Exit(0)
		default:
			continue
		}
	}

}
