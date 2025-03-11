package main

import "fmt"

// Многострочный комментарий
func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	j := 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	m := 0
	for m < 10 {
		fmt.Println(m)
		m++
	}

	users := [10]int{}

	for _, user := range users {
		fmt.Println(user)
	}

	var u = [5]string{"Tom", "John", "Julia", "Katya", "Senya"}
	for i := len(u) - 1; i >= 0; i-- {
		fmt.Println(u[i])
	}
}
