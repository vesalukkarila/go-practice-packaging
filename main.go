package main

import (
	"fmt"
	"log"

	"github.com/vesalukkarila/go-practice-packaging/greet"
)

func main() {
	greeting, err := greet.Greet("Vesa")
	if err != nil {
		log.Panic(err)

	}
	fmt.Println(greeting)
}
