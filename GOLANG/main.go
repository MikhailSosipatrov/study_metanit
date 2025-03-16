package main

import "fmt"

func main() {
	var x = 4
	var p *int
	p = &x
	fmt.Println(*p)
	fmt.Println(p)

	y := 8
	pp := &y
	fmt.Println(*pp)

	var ppp *float64
	if ppp != nil {
		fmt.Println(*ppp)
	}

	pppp := new(int)
	fmt.Println(*pppp)
	*pppp = 8
	fmt.Println(*pppp)
}
