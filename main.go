package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/vesalukkarila/go-practice-packaging/greet"
)

func main() {

	name, err := getName(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	greeting, err := greet.Greet(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	fmt.Println(greeting)
}

func getName(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("name argument not provided")

	}

	return args[1], nil

}
