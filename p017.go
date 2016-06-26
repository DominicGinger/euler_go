package main

import (
	"fmt"
)

const hundred = 7
const one_thousand = 11
const and = 3

var single = [...]int{3, 3, 5, 4, 4, 3, 5, 5, 4, 3}
var teen = [...]int{6, 6, 8, 8, 7, 7, 9, 8, 8}
var tens = [...]int{3, 6, 6, 5, 5, 5, 7, 6, 6}

func letters_in(i int) int {
	sum := 0

	if i == 0 {
		return 0
	} else if i <= 10 {
		return single[i-1]
	} else if i < 20 {
		return teen[i-11]
	} else if i < 100 {
		sum += tens[(i/10)-1]
		sum += letters_in(i % 10)
	} else if i < 1000 {
		sum += single[(i/100)-1]
		sum += hundred

		ix := letters_in(i % 100)
		if ix > 0 {
			sum += and + ix
		}
	} else {
		sum += one_thousand
	}
	return sum
}

func main() {
	total := 0
	for i := 1; i <= 1000; i++ {
		total += letters_in(i)
	}
	fmt.Println(total)
}
