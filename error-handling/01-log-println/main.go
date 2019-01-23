package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}

// log.Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
