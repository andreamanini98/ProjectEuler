package main

import (
	"fmt"
)

func main() {
	n := 1
	d := 1
	maxCycle := 0

	for i := 1; i < 1000; i++ {
		if isRec, c := isRecurring(n, i); isRec && maxCycle < c {
			maxCycle = c
			d = i
		}
	}
	fmt.Printf("Maximum cycle given by: %v\nCycle length: %v", d, maxCycle)
}

func isRecurring(n, i int) (bool, int) {
	cycle := make([]int, 0)

	defer fmt.Println()

	var q, r int = 0, n % i
	num := r * 10
	for {
		q = num / i
		r = num % i

		fmt.Print(q)

		switch {
		case r == 0:
			return false, len(cycle)
		case isInSlice(r, cycle):
			return true, len(cycle)
		default:
			cycle = append(cycle, r)
			num = r * 10
		}
	}
}

func isInSlice(i int, s []int) bool {
	for _, j := range s {
		if i == j {
			return true
		}
	}
	return false
}
