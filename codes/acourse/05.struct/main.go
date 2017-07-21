package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func (p *person) mutate() {
	p.Name = "Hacker"
	fmt.Println("In", *p)
}

func main() {
	p1 := person{
		Name:     "TT",
		NickName: "YoProgrammer",
	}
	fmt.Println(p1)

	mutatePerson(p1)
	fmt.Println("Out", p1)

	p1.mutate()
	fmt.Println("Out", p1)

	mutatePerson(p1)
	fmt.Println(p1)
}

func mutatePerson(p person) {
	p.Name = "Normal"
	fmt.Println("In:", p)
}
