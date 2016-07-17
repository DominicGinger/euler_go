package main

func dividedByEach(input int) int {
	limit := input
Out:
	for ; ; input += limit {
		for i := limit; i >= 2; i-- {
			if input%i != 0 {
				break
			}
			if i == 2 {
				break Out
			}
		}
	}
	return input
}

func p005() int {
	return dividedByEach(20)
}
