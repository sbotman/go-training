// All material is licensed under the GNU Free Documentation License

// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/disney/go-training/05-packaging/example2/counters"
)

// main is the entry point for the application.
func main() {
	// ./example2.go:15: cannot refer to unexported name counters.alertCounter
	// ./example2.go:15: undefined: counters.alertCounter
	counter := counters.alertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
