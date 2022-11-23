package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println(1)
	}()

	fmt.Println(0)
	time.Sleep(time.Second * 2)
}
