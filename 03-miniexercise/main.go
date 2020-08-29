package main

import (
	"fmt"
)

// Remember main () {content}. Be sure to keep those 
// parentheses closed before the curly brace.

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}