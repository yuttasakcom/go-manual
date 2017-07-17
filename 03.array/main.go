package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 10
	a[2] = 30
	a[3] = 40

	fmt.Println(a)
	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}

	for _, v := range a {
		fmt.Print(v)
	}

	var b [2][3]int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = j
		}
	}

	fmt.Println(b)

	c := make([]int, 5)
	c[0] = 10
	c[2] = 30
	c[3] = 40
	c = append(c, 90)

	fmt.Println(c)

	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(d)
	fmt.Println(d[2:4])
	fmt.Println(d[:4])
	fmt.Println(d[4:])
	fmt.Println(d[:])

	e := a[:2]
	e[0] = 20

	fmt.Println(a)
	fmt.Println(e)

	// make
	// f := make(map[string]string)
	// f["hello"] = "gopher"
	// f["name"] = "YoProgrammer"

	//map
	f := map[string]string{
		"hello": "gopher",
		"name":  "Yoprogrammer",
	}

	fmt.Println(f)
	fmt.Println(f["x"])

	x, ok := f["x"]
	fmt.Println(ok)
	fmt.Println(x)

	if x, ok := f["x"]; ok {
		fmt.Println(x)
	}

	for key, value := range f {
		fmt.Println(key, ":", value)
	}
}
