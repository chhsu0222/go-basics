package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrMath is the error for trying get the square root of
// negative numbers.
var ErrMath = errors.New("square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMath
	}
	return 42, nil
}
