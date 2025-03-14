package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 3,
	}
	fmt.Println(people)

	fmt.Println(people["Tom"])

	fmt.Println(people["Bob"])

	people["Alice"] = 3

	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}

	for key, value := range people {
		fmt.Println(key, value)
	}

	var orders map[string]int = make(map[string]int)
	fmt.Println(orders)

	people["Kate"] = 128
	fmt.Println(people["Kate"])

	delete(people, "Kate")
	fmt.Println(people["Kate"])
}
