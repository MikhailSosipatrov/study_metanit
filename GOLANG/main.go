package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}

func (car Car) move() {
	fmt.Println("Car moved")
}

func (aircar Aircraft) move() {
	fmt.Println("Aircraft moved")
}

func main() {
	tesla := Car{}
	boeing := Aircraft{}
	drive(tesla)
	drive(boeing)
}
