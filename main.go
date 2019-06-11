package main

import (
	"fmt"
	"os"
)

func main() {
	// var e = Entry{"zoomer.org", "kevin", "bazinga"}
	f, err := openOrCreate("test")
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(1)
	}
	entries := loadFile(f)
	for _, entry := range(entries) {
		fmt.Println(entry)
	}
}
