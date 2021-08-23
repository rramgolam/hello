package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting messagea
	names := []string{"Rishi", "Bobby", "Tommy"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// if no error
	fmt.Println(message)
}