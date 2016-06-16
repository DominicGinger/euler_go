package main

import (
	"fmt"
)

func sumOfSquares(number int) int {
	total := 0
	for ; number > 0; number-- {
		total += number * number
	}
	return total
}

func squareOfSums(number int) int {
	sum := 0
	for ; number > 0; number-- {
		sum += number
	}
	return sum * sum
}

func main() {
	su := sumOfSquares(100)
	sq := squareOfSums(100)
	fmt.Println(sq - su)
}
