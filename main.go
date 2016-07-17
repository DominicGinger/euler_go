package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var funcs = [...]func() int{
	p001, p002, p003, p004, p005,
	p006, p007, p008, p009, p010,
	p011, p012, p013, p014, p015,
	p016, p017, p018, p019, p020,
	p021, p022, p023,
}

func pad(idx int) string {
	return fmt.Sprintf("p%03d", idx+1)
}

func run(idx int) {
	start := time.Now()
	answer := funcs[idx]()
	elapsed := time.Since(start)
	fmt.Printf("Answer to %v: %v - elapsed: %v\n", pad(idx), answer, elapsed)
	if elapsed > time.Second {
		fmt.Printf("\nERROR: %v is too slow, elapsed: %v\n\n", pad(idx), elapsed)
	}
}

func runAll() {
	for i := 1; i < len(funcs); i++ {
		run(i)
	}
}

func main() {
	if len(os.Args) > 1 {
		input, _ := strconv.Atoi(os.Args[1])
		run(input - 1)
	} else {
		runAll()
	}
}
