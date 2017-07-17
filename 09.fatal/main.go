package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start...")
	doSafeWork()
	// panic("fatal error, program can not run")
	// log.Fatal("fatal error, program can not run")
	fmt.Println("Done")
}

func doFailWork() {
	panic("fail")
}

func doSafeWork() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("work fail:", r)
		}
	}()
	doFailWork()
	fmt.Println("work success")
}
