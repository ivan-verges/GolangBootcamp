package main

/*
Given a list of strings find all strings that contain a given substring.
You should split the list into chunks and process them in parallel.
Each goroutine should write the response to an output channel which will be consumed in the end to consolidate the response.
*/

//Leer Wait Group, y Resolver Ejercicio

import (
	"fmt"
	"strings"
	"sync"
)

func ContainsString(data []string, sub string, channel chan string) {
	for _, val := range data {
		if strings.Contains(val, sub) {
			channel <- val
		}
	}
}

func main() {
	data := []string{"ABCD", "CDFG", "EHGF", "FFHG", "HGBA", "ABDE"}
	target := "AB"

	chunk_count := 2

	wg := sync.WaitGroup{}

	channel := make(chan string)

	//Calls the GoRoutines with Data Chunks
	for index := 0; index < len(data); index += chunk_count {
		wg.Add(1) //Adds 1 to the WorkGroup Counter (1 Go Routine to Wait for)
		chunk := data[index:(index + chunk_count)]

		go func() {
			defer wg.Done()
			ContainsString(chunk, target, channel)
		}()
	}

	//Routine to Wait for WorkGroup and Close the Channel
	go func() {
		wg.Wait()
		close(channel)
	}()

	result := []string{}
	for val := range channel {
		result = append(result, val)
	}

	fmt.Println(result)
}

//Instruccion Select
