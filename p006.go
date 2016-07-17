package main

func sumOfSquares(number int) int {
	total := 0
	for ; number > 0; number-- {
		total += number * number
	}
	return total
}

func squareOfSums(number int) int {
	sum := 0
	for ; number > 0; number-- {
		sum += number
	}
	return sum * sum
}

func p006() int {
	su := sumOfSquares(100)
	sq := squareOfSums(100)
	return sq - su
}
