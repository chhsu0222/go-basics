package main

import "fmt"

/*
  By default channels are unbuffered, meaning that they will only accept sends (chan <-)
  if there is a corresponding receive (<- chan) ready to receive the sent value.
*/

func main() {
	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}
