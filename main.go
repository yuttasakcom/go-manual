package main

import (
	"fmt"

	"github.com/yuttasakcom/go-manual/eng"
	"github.com/yuttasakcom/go-manual/th"
)

// Hello : Hello
func Hello() string {
	return eng.Greet()
}

// ThHello : สวัสดี-Hello
func ThHello() string {
	return th.Greet()
}

func main() {
	fmt.Println(Hello())
	fmt.Println(ThHello())
}
