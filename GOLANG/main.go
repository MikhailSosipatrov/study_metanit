package main

import (
	"fmt"
)

// Многострочный комментарий
func main() {
	fmt.Println("Hello World")
	//Однострочный комментарий
	var hello string = "Hello World"
	fmt.Println(hello)

	var a, b, c string
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var (
		name string = "Tom"
		age  int    = 27
	)

	fmt.Println(name)
	fmt.Println(age)

	short_name := "Mike"
	fmt.Println(short_name)
}
