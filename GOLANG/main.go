package main

import "fmt"

func main() {
	x := 5
	addValue(x)
	fmt.Println(x)
	addValueP(&x)
	fmt.Println(x)
}

func addValue(v int) {
	v += 5
}

func addValueP(v *int) {
	*v += 5
}
