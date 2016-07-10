package main

import (
	"fmt"
	"time"
)

const size int = 28123

func isAbundant(number int) bool {
	return DivisorsSum(number) > number
}

func DivisorsSum(input int) int {
	total := 1
	for compare, upto := 2, 1; compare < input/upto; compare++ {
		if input%compare == 0 {
			total += compare
			if (input / compare) != compare {
				total += (input / compare)
			}
		}
		upto++
	}

	return total
}

func getAbundantNumbers() []int {
	abundant_numbers := []int{}

	for i := 1; i <= size; i++ {
		if isAbundant(i) {
			abundant_numbers = append(abundant_numbers, i)
		}
	}

	return abundant_numbers
}

func getAbundantSums(abundant_numbers []int) []int {
	sums := []int{}
	length := len(abundant_numbers)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum := abundant_numbers[i] + abundant_numbers[j]
			if sum <= size {
				sums = append(sums, sum)
			} else {
				break
			}
		}
	}
	return removeDuplicates(sums)
}

func removeDuplicates(sums []int) []int {
	seen := map[int]bool{}
	uniq := []int{}
	for _, v := range sums {
		if seen[v] != true {
			seen[v] = true
			uniq = append(uniq, v)
		}
	}
	return uniq
}

func getTotalWithoutSums(sums []int) int {
	total := 0
	for i := 1; i <= size; i++ {
		total += i
	}
	sums_total := 0
	for _, v := range sums {
		sums_total += v
	}
	return total - sums_total
}

func main() {
	start := time.Now()
	abundant_numbers := getAbundantNumbers()
	e1 := time.Since(start)
	abundant_sums := getAbundantSums(abundant_numbers)
	e2 := time.Since(start)
	total := getTotalWithoutSums(abundant_sums)
	e3 := time.Since(start)
	fmt.Println(total, e1, e2, e3)
}
