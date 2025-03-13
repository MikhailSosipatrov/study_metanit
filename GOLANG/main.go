package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}
