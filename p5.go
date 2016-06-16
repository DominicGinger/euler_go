package main

import (
	"fmt"
	"time"
)

func dividedByEach(input int) int {
	limit := input
Out:
	for ; ; input++ {
		for i := limit; i >= 2; i-- {
			if input%i != 0 {
				break
			}
			if i == 2 {
				break Out
			}
		}
	}
	return input
}

func main() {
	start := time.Now()
	fmt.Println(dividedByEach(20))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
