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
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions are not run.")
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
