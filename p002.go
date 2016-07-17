package main

func fibon(s1, s2, limit int) []int {
	arr := make([]int, 2)
	arr[0] = s1
	arr[1] = s2
	sum := s1 + s2
	for i := 2; sum < limit; i++ {
		sum = arr[i-1] + arr[i-2]
		arr = append(arr, sum)
	}
	return arr
}

func sumEven(arr []int) int {
	total := 0
	for _, element := range arr {
		if element%2 == 0 {
			total += element
		}
	}
	return total
}

func p002() int {
	arr := fibon(1, 2, 4000000)
	total := sumEven(arr)
	return total
}
