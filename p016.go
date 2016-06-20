package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func sumPower(input float64) int64 {
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
	return sum
}

func main() {
	start := time.Now()
	sum := sumPower(1000)
	elapsed := time.Since(start)
	fmt.Printf("answer: %d, elapsed: %s\n", sum, elapsed)
}
