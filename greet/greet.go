package greet

import (
	"errors"
	"fmt"
)

// Greet function returns greeting for given name
func Greet(name string) (string, error) {

	var greeting string
	var err error

	if len(name) < 1 {
		err = errors.New("Name must be more than 0 characters long")
		return greeting, err
	}

	greeting = fmt.Sprintf("hello %s", name)

	return greeting, err
}
