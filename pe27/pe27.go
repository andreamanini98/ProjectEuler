package main

import (
	"fmt"
)

func main() {
	max_a, max_b, max_l := 0, 0, 0

	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			l := seqPrimes(a, b)
			if max_l < l {
				max_a = a
				max_b = b
				max_l = l
			}
		}
	}
	fmt.Println("Max a, b, l, a*b:", max_a, max_b, max_l, max_a*max_b)
}

func seqPrimes(a, b int) int {
	n := 0
	seq := 0

	for {
		if !isPrime(n*n + a*n + b) {
			return seq
		}
		seq++
		n++
	}
}

func isPrime(x int) bool {
	switch {
	case x < 2:
		return false
	case x == 2:
		return true
	case x%2 == 0:
		return false
	default:
		for i := 3; i <= x/i; i += 2 {
			if x%i == 0 {
				return false
			}
		}
	}
	return true
}
