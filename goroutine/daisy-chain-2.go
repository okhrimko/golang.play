package main

import (
	"fmt"
)

func f1(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}

func main9() {
	const n = 10000

	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f1(left, right)
		left = right
	}
	go func(c chan<- int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
