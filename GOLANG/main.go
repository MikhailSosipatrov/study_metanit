package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	psn := person{"Tom", 19}
	fmt.Println(psn)
	psn.changeAge(18)
	fmt.Println(psn)
}
