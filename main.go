package main

import (
	// "fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	var e = Entry{"zoomer.org", "kevin", "bazinga"}

	entries := loadFile(fileName)

	DeleteEntry(fileName, entries, e.Site)
}
