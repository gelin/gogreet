package greet

import "fmt"

func GreetingFor(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
