package main

import (
	"fmt"
	"time"
)

func main() {
	defer time.Sleep(time.Second)
	workChan := make(chan string)
	go worker(workChan)
	workChan <- "Acoshift"
	workChan <- "Hacker"
}

func worker(ch chan string) {
	for {
		name := <-ch
		fmt.Println("Hello,", name)
	}
}
