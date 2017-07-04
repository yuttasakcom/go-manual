package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func main() {
	// p1 := person{
	// 	Name:     "TT",
	// 	NickName: "YoProgrammer",
	// }
	var p1 person
	p1.Name = "TT"
	p1.NickName = "YoProgrammer"
	fmt.Println(p1)

	// mutatePerson(p1)
	mutatePerson(&p1)
	p1.mutate()
	fmt.Println(p1)
}

//func mutatePerson(p person) {
func mutatePerson(p *person) {
	p.Name = "Hacker"
	fmt.Println("inside mutate:", p)
}

func (p *person) mutate() {
	p.Name = "Hacker"
	fmt.Println("inside mutate method:", p)
}
