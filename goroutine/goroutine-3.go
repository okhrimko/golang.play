package main

import (
	"fmt"
	"time"
)

func main3() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("do work")
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)

	done <- true
}
