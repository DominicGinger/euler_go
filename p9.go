package main

import (
	"fmt"
)

func sq(a int) int {
	return a * a
}

func productHighestPythag(input int) int {
	limit := input / 2

	for a := 1; a <= limit; a++ {
		for b := 1; b <= limit; b++ {
			for c := 1; c <= limit; c++ {
				sum := a + b + c
				if sum == input && sq(a)+sq(b) == sq(c) {
					return a * b * c
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(productHighestPythag(1000))
}
