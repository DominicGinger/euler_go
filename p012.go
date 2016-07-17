package main

func triangle_num(idx int) int {
	triangle := 0
	for i := idx; i > 0; i-- {
		triangle += i
	}
	return triangle
}

func divisors(number int) int {
	count := 0
	for i := 1; i*i <= number; i++ {
		if number%i == 0 {
			count += 2
		}
	}

	return count
}

func more_than(limit int) int {
	for i := limit; ; i++ {
		number := triangle_num(i)
		divs := divisors(number)
		if divs > limit {
			return number
		}
	}
}

func p012() int {
	return more_than(500)
}
