package greet

import (
	"fmt"
	"testing"
)

func TestGreetingFor(t *testing.T) {
	result := GreetingFor("World")
	fmt.Println(result)
	if result != "Hello, World!" {
		t.Fatal(result)
	}
}
