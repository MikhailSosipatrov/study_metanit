package main

import "fmt"

type person struct {
	name        string
	age         int
	contactInfo contactInfo
}

type contactInfo struct {
	email  string
	number string
}

func main() {
	prsn := person{
		name: "Tom",
		age:  19,
		contactInfo: contactInfo{
			email:  "email",
			number: "number",
		},
	}

	fmt.Println(prsn)
}
