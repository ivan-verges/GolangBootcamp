package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello World")
	}()

	time.Sleep(time.Second) // Wait a Second to Allow the Go Routine Run Before Main Routine Finish.
	fmt.Println("Main Function")
}
