package main

import (
	"fmt"
)

func main() {
	for a := 65; a <= 90; a++ {
		fmt.Printf("%#U\n%#U\n%#U\n\n", a, a, a)
	}
}
