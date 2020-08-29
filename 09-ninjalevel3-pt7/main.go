// Create a program that shows the “if statement” in action.
// Building on the previous hands-on exercise, (I'm merging the two) create a program that uses “else if” and “else”.

package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 4
	c := 3
	if a > b {
		fmt.Println("a is greater than at least b")
	} else if b > c {
		fmt.Println("b is greater than at least c")
	} else if c > a {
		fmt.Println("c is greater than at least a")		
	} else {
	fmt.Println("all 3 values are equivalent")
	}
}
