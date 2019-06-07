package main

import (
	"fmt"
	"os"
	// "encoding/json"
)
func main() {
	var e = Entry{"zoomer.org", "kevin", "bazinga"}
	f,err:= newDatabase("test")
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(1)
	}
	fmt.Println(entryExists(f, e))
}