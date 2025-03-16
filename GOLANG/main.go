package main

type BinaryOp func(int, int) int

type mile = uint
type kilometer = uint

type dollar uint
type ruble uint

func process(x int, y int, op BinaryOp) int {
	return op(x, y)
}

func main() {

}
