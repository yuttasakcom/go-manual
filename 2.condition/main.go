package main

import "fmt"

func main() {
	fmt.Print("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("meh!")
		return
	}

	switch fruit {
	case "1":
		fmt.Println("One")
	case "2":
		fmt.Println("Two")
	default:
		fmt.Println("--")
	}
}
