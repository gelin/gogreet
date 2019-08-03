# gogreet

Example usage of go modules.
Just a hello world library in Go.

## Usage

```go
import "github.com/gelin/gogreet"

func main() {
	fmt.Println(greet.GreetingFor("World"))
}
```

You should add to your `go.mod`:

```
require github.com/gelin/gogreet v0.0.1
```

