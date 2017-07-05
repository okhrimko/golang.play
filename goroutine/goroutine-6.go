package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main6() {

	c := multiplexer2(generator2("console <-"), generator2("file <-"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func generator2(msg string) chan Message {
	c := make(chan Message)
	waitForAll := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForAll}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForAll
		}
	}()
	return c
}

func multiplexer2(ch1, ch2 chan Message) chan Message {
	c := make(chan Message)
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
