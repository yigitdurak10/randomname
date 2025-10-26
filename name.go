package main

import (
	"log"
	"math/rand"

	"github.com/tcnksm/go-input"
)

var persons []string

func selectPC() {
	ui := &input.UI{}
	query := "a name enter"
	name, err := ui.Ask(query, &input.Options{

		Default:     "not a name",
		HideDefault: true,
		HideOrder:   true,
		Required:    true,
		Loop:        true,
	})

	if err != nil {
		log.Fatal(err)
	}

	persons = append(persons, name)

}

func main() {
	log.SetFlags(0)
	sum := 1
	for sum < 5 {
		selectPC()
		sum += 1
	}

	log.Printf("my choice: %s", persons[rand.Intn(len(persons))])
}
