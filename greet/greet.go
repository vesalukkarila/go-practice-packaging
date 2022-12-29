package greet

import (
	"fmt"
	"log"
)

// Greet function returns greeting for given name
func Greet(name string) string {
	if len(name) < 1 {
		log.Panic("Name must be more than 0 characters long")
	}
	return fmt.Sprintf("hello %s", name)
}
