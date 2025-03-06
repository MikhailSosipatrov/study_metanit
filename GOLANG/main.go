package main

import "fmt"

// Многострочный комментарий
func main() {
	var a bool = true
	var b bool = false
	var c bool = false

	fmt.Println(a || b)
	fmt.Println(c && a)
	fmt.Println(!c)
}
