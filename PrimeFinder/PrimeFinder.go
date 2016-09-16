// primeFinder
package main

import (
	"fmt"
)

func Generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func primeFinder(ch chan int, ch1 chan int) {
	var prime = <-ch
	fmt.Println(prime)
	for {
		n := <-ch
		if n%prime == 0 {
			ch1 <- n
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		ch1 := make(chan int)
		go primeFinder(ch, ch1)
		ch = ch1
	}
	fmt.Println(<-ch)
}
