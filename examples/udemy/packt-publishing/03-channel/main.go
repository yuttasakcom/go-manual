package main

import "fmt"

func main() {
	msgQueue := make(chan string, 3)
	msgQueue <- "one"
	msgQueue <- "two"
	msgQueue <- "three"

	close(msgQueue)

	for m := range msgQueue {
		fmt.Println(m)
	}
}
