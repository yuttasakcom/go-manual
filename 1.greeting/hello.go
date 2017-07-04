package main

import (
	"fmt"
	"log"
)

func main() {
	var gopher string
	gopher = "Gopher"
	fmt.Println("Hello, Gopher.")
	log.Println("Hello, Gopher")
	fmt.Printf("Hello, %s.\n", gopher)

	var name = "YoProgrammer"
	fmt.Println(name)

	nickName := "Yo"
	fmt.Printf("%s\n", nickName)

	trick := "Comment Print"
	// fmt.Println(trick)
	_ = trick
}
