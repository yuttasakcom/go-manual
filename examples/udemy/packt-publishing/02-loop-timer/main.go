package main

import (
	"fmt"
	"time"
)

func main() {
	loopTimer := time.NewTimer(time.Second * 3)

	for {
		fmt.Println("Inside the infinite loop!")

		<-loopTimer.C
		break
	}
}
