package main

import (
	"fmt"
)

func check_mod(i, v int) int {
	if i%v == 0 {
		return i
	} else {
		return 0
	}
}

func main() {
	total := 0
	for i := 1; i < 1000; i++ {
		first := check_mod(i, 3)
		total += first
		if first == 0 {
			total += check_mod(i, 5)
		}
	}
	fmt.Println(total)
}
