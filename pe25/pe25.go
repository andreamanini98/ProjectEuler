package main

import (
	"fmt"
	"math/big"
	"strings"
)

// To solve the challenge, set this to 1000 and only execute the first goroutine.
const D int = 8

func main() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	go func() {
		bf := bigFibo()
		var num *big.Int
		var idx int

		for {
			num, idx = bf()
			if countDigits(num) == D {
				break
			}
		}
		fmt.Println("bigFibo():", num)
		fmt.Println("bigFibo():", idx)
		close(done1)
	}()

	go func() {
		var num *big.Int
		var idx int = 1

		for {
			num = fibonacci(idx)
			if countDigits(num) == D {
				break
			}
			idx++
		}
		fmt.Println("fibonacci():", num)
		fmt.Println("fibonacci():", idx)
		close(done2)
	}()
	<-done1
	<-done2
}

func bigFibo() func() (*big.Int, int) {
	var x, y *big.Int = big.NewInt(1), big.NewInt(1)
	var idx int = 0

	return func() (*big.Int, int) {
		res := x
		x = y
		y = new(big.Int).Add(res, y)
		idx++
		return res, idx
	}
}

func fibonacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	return new(big.Int).Add(fibonacci(n-1), fibonacci(n-2))
}

func countDigits(n *big.Int) int {
	return len(strings.Split(n.String(), ""))
}
