package main

import "strconv"

func reverse(str string) string {
	length := len(str)
	runes := make([]rune, length)
	for _, rune := range str {
		length--
		runes[length] = rune
	}
	return string(runes)
}

func isPalindrone(number int) bool {
	str := strconv.Itoa(number)
	if len(str)%2 != 0 {
		return false
	}
	midway := len(str) / 2
	first := str[0:midway]
	second := reverse(str[midway:len(str)])

	return first == second
}

func biggestPalindroneFactor(start, end int) int {
	biggest := 0
	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			factor := i * j
			if isPalindrone(factor) && factor > biggest {
				biggest = factor
			}
		}
	}
	return biggest
}

func p004() int {
	biggest := biggestPalindroneFactor(100, 999)
	return biggest
}
