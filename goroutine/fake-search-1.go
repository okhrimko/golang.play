package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web1   = fakeSearch("web1")
	web2   = fakeSearch("web2")
	image1 = fakeSearch("image1")
	image2 = fakeSearch("image2")
	video1 = fakeSearch("video1")
	video2 = fakeSearch("video2")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := google("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

	/*
		rand.Seed(time.Now().UnixNano())
		start := time.Now()
		result := First("golang",
			fakeSearch("replica 1"),
			fakeSearch("replica 2"))
		elapsed := time.Since(start)
		fmt.Println(result)
		fmt.Println(elapsed)
	*/
}

type Search func(query string) Result

type Result string

/* solution 1
func Google(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return results
}
*/

/* solution 2
func Google(query string) (results []Result) {
	ch := make(chan Result)
	go func() {
		ch <- web(query)
	}()
	go func() {
		ch <- image(query)
	}()
	go func() {
		ch <- video(query)
	}()

	timeout := time.After(2 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return results
}
*/

func google(query string) (results []Result) {
	ch := make(chan Result)
	go func() {
		ch <- first(query, web1, web2)
	}()
	go func() {
		ch <- first(query, image1, image2)
	}()
	go func() {
		ch <- first(query, video1, video2)
	}()

	timeout := time.After(2 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return results
}

func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
