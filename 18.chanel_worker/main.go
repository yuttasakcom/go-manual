package main

import "fmt"

func main() {
	workChan := make(chan string)
	go worker(workChan)
	workChan <- "Yo"
	workChan <- "Hacker"
}

func worker(ch chan string) {
	for {
		name := <-ch
		fmt.Println("Hello,", name)
	}
}
