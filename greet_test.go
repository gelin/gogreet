package greet

import "testing"

func TestGreetingFor(t *testing.T) {
	result := GreetingFor("World")
	if result != "Hello, World!" {
		t.Fatal(result)
	}
}
