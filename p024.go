package main

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}

func p024() int {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 999999
	seq := []int{}
	seq = append(seq, n/fact(9))
	for i := 8; i >= 0; i-- {
		v := i
		if i != 0 {
			n %= fact(i + 1)
			v = n / fact(i)
		}
		seq = append(seq, v)
	}

	r := 0
	for i, v := range seq {
		x := 1
		for j := len(seq) - 1; j > i; j-- {
			x *= 10
		}
		r += input[v] * x
		input = append(input[:v], input[v+1:]...)
	}

	return r
}
