package main

import (
	"fmt"
	"time"
)

func main4() {

	c := make(chan bool)

	go func() {
		// here executing job stuff
		time.Sleep(10 * time.Second)
		c <- true
	}()

	select {
	case res := <-c:
		fmt.Println("result: ", res)
	case <-time.After(11 * time.Second):
		fmt.Println("timeout")
	}

	//http.HandleFunc
}
