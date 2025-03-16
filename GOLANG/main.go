package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{"Tom", 25}
	fmt.Println(tom)
	fmt.Println(tom.age)
	fmt.Println(tom.name)

	tomP := &tom
	fmt.Println(tomP)
	fmt.Println(tomP.age)
	fmt.Println(tomP.name)
	fmt.Println(((*tomP).age))
}
