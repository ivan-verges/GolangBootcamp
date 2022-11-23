package main

import (
	"fmt"
	"sync"
	"time"
)

type MyData struct {
	Value int
	Mutex sync.Mutex
}

func AddValue(data *MyData) {
	data.Mutex.Lock()

	fmt.Println("Current Value:", data.Value)
	data.Value++

	data.Mutex.Unlock()
}

func main() {
	data := MyData{Value: 0}

	for index := 0; index < 1000; index++ {
		go AddValue(&data)
	}

	time.Sleep(time.Second)

	fmt.Println("Final Value:", data.Value)
}
