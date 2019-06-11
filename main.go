package main

import (
	// "bufio"
	// "fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	var e = Entry{"zoomer.org", "kevin", "bazinga"}

	entries := loadFile(fileName)

	UpdateEntry(fileName, entries, e.Site)
}
