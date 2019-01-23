package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)

	// block until channel closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
