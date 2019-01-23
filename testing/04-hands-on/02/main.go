package main

import (
	"fmt"

	"github.com/chhsu0222/go-basics/testing/04-hands-on/02/quote"
	"github.com/chhsu0222/go-basics/testing/04-hands-on/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
