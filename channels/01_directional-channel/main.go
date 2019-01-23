package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	receive(c) // block
	fmt.Println("About to exit")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
