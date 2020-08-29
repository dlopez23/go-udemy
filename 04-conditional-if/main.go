package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}

	// scope

	if x := 42; x == 2 {
		fmt.Println("1")
	}
	fmt.Println("here's a statement")
	fmt.Println("something else")

	if x := 42; x == 42 {
		fmt.Println("1")
	}
	fmt.Println("here's a statement")
	fmt.Println("something else")

	// if / else if / else if/ ... / else
	// i believe it has to end in else otherwise it's a loop?

	x := 11
	if x == 10 {
		fmt.Println("our value was 10")
	} else if x == 11 {
		fmt.Println("our value was 11")
	} else if x == 12 {
		fmt.Println("our value was 12")
	} else {
		fmt.Println("our value was not 10, 11, or 12")
	}
}
