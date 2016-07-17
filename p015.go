package main

func factorial(input float64) float64 {
	output := float64(1)
	for ; input > 0; input-- {
		output *= input
	}
	return output
}

func p015() int {
	var width float64 = 20
	top := factorial(width + width)
	bottom := factorial(width) * factorial(width)
	answer := top / bottom
	return int(answer)
}

// https://en.wikipedia.org/wiki/Permutation#Permutations_of_multisets
