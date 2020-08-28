package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for a := 0; a <= 10; a++ {
		fmt.Printf("The outer loop: %d\n", a)
		for b := 0; b < 3; b++ {
			fmt.Printf("\t\t The inner loop: %d\n", b)
		}
	}
}
