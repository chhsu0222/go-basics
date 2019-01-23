package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("Deferred functions are run in the usual way.")
}

// Panicln is equivalent to Println() followed by a call to panic().
