package main

func highestPrimeFactor(input int) int {
	output := 1
	for output%2 == 0 {
		output = 2
		input = input / 2
	}

	compare := 3
	for ; input > 1; compare++ {
		for input%compare == 0 {
			output = compare
			input = input / compare
		}
	}
	return output
}

func p003() int {
	output := highestPrimeFactor(600851475143)
	return output
}
