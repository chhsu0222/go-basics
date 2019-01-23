package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	// block
	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to exit")

}

// send channel
func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// receive channel
func receive(e, o <-chan int, fin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fin)
}
