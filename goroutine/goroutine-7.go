package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := multiplexer3(generator3("console <-"), generator3("file <-"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm leaving")
}

func generator3(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func multiplexer3(ch1, ch2 chan string) chan string {
	c := make(chan string, 2)

	go func() {
		for {
			select {
			case s := <-ch1:
				c <- s
			case s := <-ch2:
				c <- s
			}
		}
	}()

	return c
}
