package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer f2.Close()

	fmt.Println("Check the log.txt filr in the directory.")
}

// SetOutput sets the output destination for the standard logger.
