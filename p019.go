package main

import (
	"fmt"
)

var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func sunday(i int) bool {
	if i%7 == 0 {
		return true
	}
	return false
}

func main() {
	total_days := 0
	total_sundays := 0
	for i := 0; i < 101; i++ {
		if i%4 == 0 {
			months[1] = 29
		} else {
			months[1] = 28
		}
		for j := 0; j < 12; j++ {
			total_days += months[j]
			if sunday(total_days) && i > 0 {
				total_sundays++
			}
		}
	}
	fmt.Println(total_sundays)
}
