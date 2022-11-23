package main

import (
	"fmt"
)

func AddToChannel(channel chan int) {
	channel <- 5
}

func main() {
	channel := make(chan int)
	go AddToChannel(channel)

	value, ok := <-channel
	if ok {
		fmt.Println(value)
	}
}
