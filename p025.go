package main

import (
	"math/big"
)

func p025() int {
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := 1
	for a.Cmp(&limit) < 0 && b.Cmp(&limit) < 0 {
		a.Add(a, b)
		a, b = b, a
		c++
	}
	return c
}
