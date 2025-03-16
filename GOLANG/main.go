package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) printName() {
	fmt.Println("Name:", p.name)
}

func main() {
	psn := person{"Tom", 19}
	psn.printName()
}
