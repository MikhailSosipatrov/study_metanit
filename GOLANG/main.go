package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {
	result := 0
	for _, n := range numbers {
		if criteria(n) {
			result += n
		}
	}
	return result
}

func add(x int, y int) int      { return x + y }
func subtract(x int, y int) int { return x - y }
func multiply(x int, y int) int { return x * y }

func selectFunction(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func main() {
	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEven := sum(slice, isEven)
	fmt.Println(sumOfEven)

	sumOfPositives := sum(slice, isPositive)
	fmt.Println(sumOfPositives)

	f := selectFunction(1)
	fmt.Println(f(1, 2))
}
