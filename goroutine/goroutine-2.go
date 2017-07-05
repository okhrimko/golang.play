package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main2() {

	sig := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sig)

	go func() {
		sgn := <-sig
		fmt.Println(sgn)
		done <- true
	}()

	fmt.Println("awating signal")
	<-done
	fmt.Println("exit")
}
