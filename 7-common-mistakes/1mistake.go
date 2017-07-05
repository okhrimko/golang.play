package main

import (
	"fmt"
)

func main() {

	//a := make([]string, 5)
	//a := [5]string{"1", "2", "3", "4", "5"}
	/*b := a[:0]
	for ok, x := range a {
		if ok >= 0 {
			b = append(b, x)
		}
	}*/

	/*b := make([]string, len(a))
	copy(b, a[:len(a)])
	b[1] = "7"*/

	//fmt.Println([]string(nil))
	//fmt.Println(len([]string(nil)))
	/*
		a := []string(nil)
		if a == nil {
			fmt.Println("wow1")
		}
	*/

	c := new([]string)
	if *c == nil {
		fmt.Println("wow2")
	}
	f := make([]string, 5)
	if f == nil {
		fmt.Println("wow3")
	}

	// or
	//append(a, b...)
	//b := append([]string(nil), a[0:len(a)]...)

	//fmt.Println(b)

}
