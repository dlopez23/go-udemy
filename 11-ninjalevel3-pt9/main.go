// Create a program that uses a switch statement with the
// switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import ("fmt")

func main(){
	favSport := "futbol"
	switch favSport {
	case "baseball":
		fmt.Println("no")
	case "soccer":
		fmt.Println("did you mean futbol?")
	case "futbol":
		fmt.Println("simon!")
	case "basketball":
		fmt.Println("not bad but not fave")
	case "football":
		fmt.Println("wrong spelling")
	}
}