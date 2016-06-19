package main

import (
	"fmt"
)

func even(input int) int {
	return input / 2
}

func odd(input int) int {
	return (input * 3) + 1
}

func do_loop(input int) int {
	count := 1
	for input > 1 {
		count++
		if input%2 == 0 {
			input = even(input)
		} else {
			input = odd(input)
		}
	}
	return count
}

func longest_chain(limit int) int {
	biggest := 1
	output := 1
	for i := 1; i < limit; i++ {
		count := do_loop(i)
		if count > biggest {
			biggest = count
			output = i
		}
	}

	return output
}

func main() {
	fmt.Println(longest_chain(1000000))
}
