package main

import (
	"fmt"
	"time"
)

func isPrime(input int) bool {
	if input == 2 {
		return true
	}
	if input%2 == 0 {
		return false
	}

	compare := 3
	upto := 2
	for ; compare < input/upto; compare++ {
		if input%compare == 0 && input != compare {
			return false
		}
		upto++
	}

	return true
}

func sumOfPrimesBelow(input int) int {
	total := 0
	for i := 2; i < input; i++ {
		if isPrime(i) {
			total += i
		}
	}
	return total
}

func main() {
	start := time.Now()
	fmt.Println(sumOfPrimesBelow(2000000))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
