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

func countPrime(input int) int {
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == input {
			return i
		}
	}
	return 0
}

func main() {
	start := time.Now()
	answer := countPrime(10001)
	elapsed := time.Since(start)
	fmt.Printf("%d %s\n", answer, elapsed)
}
