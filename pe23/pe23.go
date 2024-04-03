package main

import (
	"fmt"
)

func main() {
	fmt.Println("SUM:", getNotAbundants())
}

func getNotAbundants() int {
	a := getIntsFromAbundantSum()
	sum := 0

	for i := 0; i < 28123; i++ {
		fmt.Println(i)
		if !isAlreadyChosen(i, a) {
			sum += i
		}
	}

	return sum
}

func getIntsFromAbundantSum() []int {
	a := getAbundants()
	abundantsSum := make([]int, 0)

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			tmp := a[i] + a[j]
			if tmp < 28123 {
				abundantsSum = append(abundantsSum, tmp)
			}
		}
	}
	fmt.Println(len(abundantsSum))
	return abundantsSum
}

func isAlreadyChosen(x int, a []int) bool {
	for _, u := range a {
		if u == x {
			return true
		}
	}
	return false
}

func getAbundants() []int {
	a := make([]int, 0)
	for i := 12; i < 28123; i++ {
		if isAbundant(i) {
			a = append(a, i)
		}
	}
	return a
}

func isAbundant(x int) bool {
	return x < getDivisorsSum(x)
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
