package main

import (
	"fmt"
	"math/big"
)

func main() {
	terms := make([]*big.Int, 0)

	fmt.Println(pow(big.NewInt(3), big.NewInt(3)))

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			p := pow(bigA, bigB)
			if !isInSlice(p, terms) {
				terms = append(terms, p)
			}
		}
	}
	fmt.Println(terms)
	fmt.Println(len(terms))
}

func pow(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(0).Mul(a, pow(a, big.NewInt(b.Int64()-1)))
}

func isInSlice(a *big.Int, s []*big.Int) bool {
	for i := 0; i < len(s); i++ {
		if a.Cmp(s[i]) == 0 {
			return true
		}
	}
	return false
}
