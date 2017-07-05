package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main8() {
	const n = 100
	//var right chan int
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
		//fmt.Println("Loop ", i)
	}
	go func(c chan int) { c <- 1 }(right)

	fmt.Println(<-leftmost)
}
