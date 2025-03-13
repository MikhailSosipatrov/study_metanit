package main

import "fmt"

func main() {
	defer finish()
	defer fmt.Println("Program is exiting")
	var x [5]int = [5]int{}
	var y = [5]int{}
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
	fmt.Println(divide(5, 1))

	var users []string = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users)
	fmt.Println(len(users))
	fmt.Println(cap(users))
	fmt.Println(users[0])

	for _, value := range users {
		fmt.Println(value)
	}

	users = append(users, "Tom")

	for _, value := range users {
		fmt.Println(value)
	}

	var myUsers = make([]string, 3)
	myUsers[0] = "Tom"
	fmt.Println(myUsers)

	myUsers = append(myUsers, "Tom")
	fmt.Println(myUsers)

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	users1 := initialUsers[2:6]
	users2 := initialUsers[2:]
	users3 := initialUsers[5:]

	fmt.Println(users1)
	fmt.Println(users2)
	fmt.Println(users3)

	fmt.Println(users1)
	usrs := append(users1[:1], users1[2:]...)
	fmt.Println(usrs)
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
