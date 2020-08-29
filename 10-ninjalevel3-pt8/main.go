// Create a program that uses a switch statement with no switch expression specified.

package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("won't print")
	case true:
		fmt.Println("will print")
	}
}

// Cases use : instead of {} at least here