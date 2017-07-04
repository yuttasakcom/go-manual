package main

import "fmt"

func main() {
	a := 10
	ptrA := &a

	fmt.Println(a)
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 20
	fmt.Println(a)
	fmt.Println(*ptrA)

	b, c := 1, 2
	r := add(b, c)
	fmt.Println(r)

	d := [5]int{}
	fmt.Println(d)
	mutateArray(d[:])
	fmt.Println(d)
}

func add(x, y int) int {
	p := 1 // variable in stack
	return x + y + p
}

// slide คือ pointer ที่ชี้อยู่ใน array เป็นส่วนหนึ่งของ array
func mutateArray(a []int) {
	a[0] = 10
	// fmt.Println(a)
}
