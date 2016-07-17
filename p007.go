package main

func isPrime(input int) bool {
	if input == 2 {
		return true
	}
	if input%2 == 0 {
		return false
	}

	for i := 3; i*i <= input; i += 2 {
		if input%i == 0 {
			return false
		}
	}
	return true
}

func countPrime(input int) int {
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == input {
			return i
		}
	}
}

func p007() int {
	return countPrime(10001)
}
