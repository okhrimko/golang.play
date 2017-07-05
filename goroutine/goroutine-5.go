package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main5() {
	/* example 1 - generator
	p := generator("print")
	w := generator("write")
	for i := 0; i < 5; i++ {
		fmt.Println(<-p)
		fmt.Println(<-w)
	}
	fmt.Println("I'm leaving")
	*/

	/* example 2 = multiplexer */
	c := multiplexer(generator("console <-"), generator("file <-"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm leaving")
}

func generator(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func multiplexer(ch1, ch2 chan string) chan string {
	c := make(chan string, 2)
	go func() {
		fmt.Println("c <- <-ch1")
		for {
			c <- <-ch1
		}
	}()

	go func() {
		fmt.Println("c <- <-ch2")
		for {
			c <- <-ch2
		}
	}()

	return c
}
