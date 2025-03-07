package main

import "fmt"

// Многострочный комментарий
func main() {
	a := 5
	b := 2

	if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}
