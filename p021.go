package main

import (
	"fmt"
)

func divisors_sum(input int) int {
	total := 0
	for i := 1; i <= input/2; i++ {
		if input%i == 0 {
			total += i
		}
	}
	return total
}

func main() {
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < 10000; i++ {
		m[i] = divisors_sum(i)
	}

	sum := 0
	for k, v := range m {
		if m[v] == k && v != k {
			sum += k
		}
	}

	fmt.Println(sum)
}
