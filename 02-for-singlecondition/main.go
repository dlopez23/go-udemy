package main

import (
	"fmt"
)

func main()  {

	// First example

	x:= 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

	// Second example

	y:=1
	for {
	if y > 9 {
		break
	}
	fmt.Println(y)
	y++
	}
	fmt.Println("Doner")	

	// Third example
	
	z:=1
	for {
		z++	
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)
	}
	fmt.Println("Donest")	


	// Third example my take

	d:=1
	for {	
	if d > 100 {
		break
	}
	if d%2 != 0 {
		d++
		continue
	}
	fmt.Println(d)
	d++
	}
	fmt.Println("Donest 2")	
}