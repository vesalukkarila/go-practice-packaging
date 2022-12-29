package greet

import "fmt"

// Greet function returns greeting for given name
func Greet(name string) string {
	return fmt.Sprintf("hello %s", name)
}
