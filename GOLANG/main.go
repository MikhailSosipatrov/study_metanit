package main

import "fmt"

// Многострочный комментарий
func main() {
	add(1, 2, 3)
	add(1, 2, 3, 4)
}

func add(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
