package main

import (
	"fmt"
	"time"
)

func highestPrimeFactor(input int64) int64 {
	output := int64(1)
	for output%2 == 0 {
		output = int64(2)
		input = input / 2
	}

	compare := int64(3)
	for ; input > 1; compare++ {
		for input%compare == 0 {
			output = compare
			input = input / compare
		}
	}
	return output
}

func main() {
	start := time.Now()
	output := int64(1)
	for i := 100000; i > 1; i-- {
		output = highestPrimeFactor(600851475143)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time took: %s\n", elapsed)
	fmt.Println(output)
}
