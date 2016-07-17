package main

var primes = []int{}

func isComposite(input int) bool {
	for _, v := range primes {
		if input%v == 0 {
			return true
		}
	}
	return false
}

func sumOfPrimesBelow(input int) int {
	total := 0
	for i := 3; i < input; i += 2 {
		// if isComposite(i) {
		// continue
		// }

		if isPrime(i) {
			primes = append(primes, i)
			total += i
		}
	}
	return total + 2
}

func p010() int {
	return sumOfPrimesBelow(2000000)
}
