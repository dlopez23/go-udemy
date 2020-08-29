package main

import (
	"fmt"
)

func main() {
	for a := 65; a <= 90; a++ {
		fmt.Println(a)
		for b := 0; b < 3; b++ {
			fmt.Printf("\t%#U\n", a)
		}
	}
}
