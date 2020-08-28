package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for a := 0; a <= 10; a++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", a, j)
		}
	}
}
