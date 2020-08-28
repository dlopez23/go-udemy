package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for a := 0; a <= 10; a++ {
		fmt.Printf("The outer loop: %d\n", a)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}
