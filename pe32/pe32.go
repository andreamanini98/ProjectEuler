package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	pan, getSum := isPandigital()

	for i := 1; i <= 999; i++ {
		for j := 1; j <= 9999; j++ {
			pan(i, j)
		}
	}
	fmt.Println(getSum())
}

func isPandigital() (func(int, int) bool, func() int) {
	var products = map[int]bool{}
	var sum int

	return func(x, y int) bool {
			p := x * y
			if _, ok := products[p]; !ok {
				digitsX := strings.Split(strconv.Itoa(x), "")
				digitsY := strings.Split(strconv.Itoa(y), "")
				digitsP := strings.Split(strconv.Itoa(p), "")

				if len(digitsX)+len(digitsY)+len(digitsP) > 9 {
					return false
				}

				allDigits := append(digitsX, append(digitsY, digitsP...)...)
				digitsCount := make([]int, 9)
				for _, v := range allDigits {
					vInt, _ := strconv.Atoi(v)
					if 1 <= vInt && vInt <= 9 {
						digitsCount[vInt-1]++
					} else {
						return false
					}
				}
				for _, v := range digitsCount {
					if v != 1 {
						return false
					}
				}
				products[p] = true
				sum += p
				return true
			}
			return false
		},
		func() int {
			return sum
		}
}
