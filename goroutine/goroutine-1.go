package main

import (
	"fmt"
)

func printChan(c chan int, done chan bool) {
	fmt.Println("start goroutine")
	fmt.Printf("read from channel: %d\n", <-c)
	fmt.Println("end goroutine")
	done <- true
}

func main1() {

	c := make(chan int)
	done := make(chan bool)
	go printChan(c, done)

	fmt.Println("Start")
	c <- 1
	<-done
	fmt.Println("End")
}
