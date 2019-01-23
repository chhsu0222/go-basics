package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even number:", v)
		case v := <-o:
			fmt.Println("Odd number:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok", i, ok)
				fmt.Println("Quit")
				return
			}

			fmt.Println("From comma ok", i, ok)
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
