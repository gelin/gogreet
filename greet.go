package greet

import (
	"fmt"
	"github.com/gelin/gogreet/format"
)

func GreetingFor(name string) string {
	return fmt.Sprintf(format.GreetingFormat(), name)
}
