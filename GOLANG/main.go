package main

import "fmt"

func main() {
	defer finish()
	defer fmt.Println("Program is exiting")
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
	fmt.Println(divide(5, 1))
	fmt.Println(divide(5, 0))
}

func finish() {
	fmt.Println("finish")
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}
