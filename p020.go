package main

import (
	"math/big"
	"strconv"
	"strings"
)

func bigFactorial(num int) *big.Int {
	f := big.NewInt(1)
	for i := 1; i <= num; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}

func sum_digits(num *big.Int) int {
	str := num.String()
	numbers := strings.Split(str, "")
	sum := 0
	for i := 0; i < len(numbers); i++ {
		val, _ := strconv.Atoi(numbers[i])
		sum += val
	}
	return sum
}

func p020() int {
	input := 100
	fact := bigFactorial(input)
	return sum_digits(fact)
}
