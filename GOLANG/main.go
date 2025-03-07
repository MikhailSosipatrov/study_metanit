package main

import "fmt"

// Многострочный комментарий
func main() {
	n := 4
	fmt.Println(n << 2)

	var numbers [5]int
	fmt.Println(numbers)

	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	numbrs := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbrs)

	var numbs = [...]int{1, 2, 3, 4}
	fmt.Println(numbs)
}
