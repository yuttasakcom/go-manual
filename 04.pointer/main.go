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

	a1 := 1
	b1 := 2

	fmt.Println(add2(&a1, &b1))

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func add(x, y int) int {
	p := 1 // variable in stack
	return x + y + p
}

func add2(a, b *int) int {
	return *a + *b
}

// slide คือ pointer ที่ชี้อยู่ใน array เป็นส่วนหนึ่งของ array
func mutateArray(a []int) {
	a[0] = 10
	// fmt.Println(a)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 5
}
