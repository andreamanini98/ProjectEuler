package main

import (
	"fmt"
)

func main() {
	amicables := make([]int, 0)

	for i := 1; i < 10000; i++ {
		a := getDivisorsSum(i)
		if areAmicable(i, a) {
			if !isAlreadyChosen(a, amicables) {
				amicables = append(amicables, a)
			}
		}
	}

	sum := 0
	for _, u := range amicables {
		sum += u
	}

	fmt.Println(sum)
}

func isAlreadyChosen(x int, a []int) bool {
	for _, u := range a {
		if u == x {
			return true
		}
	}
	return false
}

func areAmicable(x, y int) bool {
	// Call with x, y such that y = getDivisorsSum(x).
	if x == y {
		return false
	}
	return getDivisorsSum(y) == x
}

func getDivisorsSum(x int) int {
	sum := 0
	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			sum += i
		}
	}
	return sum
}
