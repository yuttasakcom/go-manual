package main

import "fmt"

type person struct {
	Name string
}

func (p person) Talk() {
	fmt.Println("Hello, I'm", p.Name)
}

type cat struct{}

func (cat) Talk() {
	fmt.Println("Nyaa nyaa")
}

type dog struct{}

func (*dog) Talk() {
	fmt.Println("Wan wan")
}

type talkable interface {
	Talk()
}

func talkWith(t talkable) {
	t.Talk()
}

func main() {
	p := person{"Gopher"}
	// p.Talk()
	talkWith(p)
	talkWith(cat{})
	talkWith(&dog{})

	checkType("Gopher")
	checkType(10)
	checkType(true)

	m := make(map[interface{}]interface{})
	m[1] = "Hello"
	m["name"] = "Gopher"
	q := person{"TT"}
	m[q] = "YoProgrammer"
	fmt.Println(m[q])

	if x, ok := m["name"].(string); ok {
		fmt.Println(x)
	}
}

// func checkType(v interface{}) {
// 	switch p := v.(type) {
// 	case string:
// 		fmt.Println("String", p)
// 	case int:
// 		fmt.Println("Int", p+10)
// 	case bool:
// 		fmt.Println("Bool", !p)
// 	}
// }

func checkType(v interface{}) {
	// p, ok := v.(string)
	// if ok {
	// 	fmt.Println(p)
	// } else {
	// 	fmt.Println("v is not string")
	// }
	p, _ := v.(string)
	fmt.Println(p)
}
