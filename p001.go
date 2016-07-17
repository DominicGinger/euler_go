package p001

func checkMod(i, v int) int {
	if i%v == 0 {
		return i
	}
	return 0
}

func P001() int {
	total := 0
	for i := 1; i < 1000; i++ {
		first := checkMod(i, 3)
		total += first
		if first == 0 {
			total += checkMod(i, 5)
		}
	}
	return total
}
