package main

import "fmt"

// Многострочный комментарий
func main() {
	x, z := add(1, 4, "Tom", "Bobs")
	fmt.Println(x, z)
}

func add(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullname = firstName + " " + lastName
	return z, fullname
}
