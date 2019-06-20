package main

import (
	// "bufio"
	"fmt"
	"os"
)

// var (
// 	fileName string
// 	key []byte
// )

func main() {
	fmt.Println("coucou")
	fileName := os.Args[1]
	key, _ := aesKey([]byte("hihihi"))
	// var e = Entry{"zoomer.org", "kevin", "bazinga"}

	entries := loadFile(fileName, key)

	CreateEntry(fileName, entries, key)
	CreateEntry(fileName, entries, key)
	ReadEntry(entries, "zoomer")
}
