package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (m mathError) Error() string {
	return fmt.Sprintf("a math error occured: %v %v %v", m.lat, m.long, m.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("square root of negative number: %v", f)
		return 0, mathError{"50.22 N", "99.45 W", me}
	}
	return 42, nil
}

// https://golang.org/pkg/builtin/#error
