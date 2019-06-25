package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileName string
	key      []byte
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: pwd <password database>")
	}
	fileName = os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("enter your password\n~> ")
	scanner.Scan()
	key = aesKey([]byte(scanner.Text()))

	entries := loadFile(fileName, key)

	for {
		fmt.Printf("Enter a command:\n")
		fmt.Printf(
			"[c]reate\t[r]ead <entry>\t[u]pdate <entry>\t[d]elete <entry>\t[l]ist all sites\t[q]uit\n")
		scanner.Scan()
		command := []string{""}
		if scanner.Text() != "" {
			lowerCase := strings.ToLower(scanner.Text())
			command = strings.Fields(lowerCase)
		}
		switch command[0] {
		case "c":
			if len(command) > 1 {
				fmt.Println("usage: c")
				continue
			}
			entries = CreateEntry(fileName, entries, key)
		case "r":
			if len(command) != 2 {
				fmt.Println("usage: r <sitename>")
				continue
			}
			ReadEntry(entries, command[1])
		case "u":
			if len(command) != 2 {
				fmt.Println("usage: u <sitename>")
				continue
			}
			entries = UpdateEntry(fileName, entries, command[1], key)
		case "d":
			if len(command) != 2 {
				fmt.Println("usage: d <sitename>")
				continue
			}
			entries = DeleteEntry(fileName, entries, command[1], key)
		case "l":
			if len(command) != 1 {
				fmt.Println("usage: l")
				continue
			}
			PrintSites(entries)
		case "q":
			os.Exit(0)
		default:
			continue
		}
	}

}
