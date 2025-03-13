package main

import "fmt"

func main() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(1, 3))
	fmt.Println(f(0, -1))
	action(1, 2, func(x int, y int) int {
		return x + y
	})

	f2 := selectFn(1)
	fmt.Println(f2(4, 5))

	f3 := square()
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
}

func action(n1, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }
	} else if n == 2 {
		return func(x, y int) int { return x - y }
	} else {
		return func(x, y int) int { return 0 }
	}
}

func square() func() int {
	x := 2

	return func() int {
		x++
		return x * x
	}
}
