package main

import (
	"context"
	"fmt"
	"time"
)

func getValuefromChannel(ch chan<- string,
	duration time.Duration) {
	time.Sleep(duration)
	ch <- "The End"
}

func somaTodos(x ...int) int {
	counter := 0
	for _, y := range x {
		counter += y
	}
	return counter
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ch1 := make(chan string)
	go getValuefromChannel(ch1, 10*time.Second)
	ch2 := make(chan string)
	go getValuefromChannel(ch2, 10*time.Second)
	ch3 := make(chan string)
	go getValuefromChannel(ch3, 10*time.Second)

	select {
	case returnValue := <-ch1:
		fmt.Println(returnValue)
	case returnValue := <-ch2:
		fmt.Println(returnValue)
	case returnValue := <-ch3:
		fmt.Println(returnValue)
	case <-ctx.Done():
		fmt.Println("Timeout")

	}
}
