package main

import (
	"fmt"
)

func main() {
	fmt.Print("Input Score: ")
	var score int

	fmt.Scanln(&score)

	if score < 50 {
		fmt.Println("Grade: 0")
	} else if score < 60 {
		fmt.Println("Grade: 1")
	} else {
		fmt.Println("Grade: S")
	}
}
