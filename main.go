package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	var e = Entry{"zoomer.org", "kevin", "bazinga"}

	entries := loadFile(fileName)
	for i := 0; i < 10; i++ {
		fmt.Println(genPW(24))
	}
	UpdateEntry(fileName, entries, e.Site)
}
