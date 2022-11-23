package main

import (
	"fmt"
)

type MusicalPlayer interface {
	Greetings()
}

type Trumpeter struct {
	Name string
}

func (musician Trumpeter) Greetings() {
	fmt.Println("Hello Everyone, I Am", musician.Name, "The Trumpeter")
}

func CreateTrumpeter(name string) Trumpeter {
	return Trumpeter{Name: name}
}

type Violinist struct {
	Name string
}

func (musician Violinist) Greetings() {
	fmt.Println("Hello Everyone, I Am", musician.Name, "The Violinist")
}

func CreateViolinist(name string) Violinist {
	return Violinist{Name: name}
}

func main() {

	musicians := []MusicalPlayer{}
	musicians = append(musicians, CreateTrumpeter("Juan"))
	musicians = append(musicians, CreateViolinist("Pedro"))

	for _, musician := range musicians {
		musician.Greetings()
	}
}
