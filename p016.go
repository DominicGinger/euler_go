package main

import (
	"math"
	"strconv"
)

func sumPower(input float64) int {
	number := math.Pow(2, input)
	str := strconv.FormatFloat(number, 'e', int(input), 64)
	str = str[:(len(str) - 4)]
	str = str[0:1] + str[2:]
	length := len(str)
	sum := int64(0)
	for i := 0; i < length; i++ {
		next, _ := strconv.ParseInt(string(str[i]), 10, 64)
		if next != 0 {
			sum += next
		}
	}
	return int(sum)
}

func p016() int {
	return sumPower(1000)
}
