package main

import "fmt"

// Многострочный комментарий
func main() {
	var a int8 = -4
	var b uint64 = 5
	var c rune = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var fl float64 = 3.14
	fmt.Println(fl)

	var comp complex64 = 1 + 2i
	fmt.Println(comp)

	var isActive bool = true
	fmt.Println(isActive)

	var name string = "Tom Soier"
	fmt.Println(name)

	first, secong := 1, "str"
	fmt.Println(first, secong)

}
