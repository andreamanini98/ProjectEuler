package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var precision float64 = 0.000000000001

// At the end you have to check the resulting values to satosfy the requirement.
// The correct ones are:
// 49 / 98 = 4 / 8
// 16 / 64 = 1 / 4
// 19 / 95 = 1 / 5
// 26 / 65 = 2 / 5

func main() {
	for i := 11; i <= 99; i++ {
		for j := i; j <= 99; j++ {
			if i != j {
				if !(i%10 == 0 || j%10 == 0) {
					if c, n, d := isCancelling(i, j, precision); c {
						fmt.Println("i:", i, ", j:", j, ", n:", n, ", d:", d)
					}
				}
			}
		}
	}
}

func isCancelling(n, d int, precision float64) (bool, int, int) {
	var frac float64 = float64(n) / float64(d)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {

			var diff float64 = math.Abs(float64(frac) - float64(i)/float64(j))
			if diff <= precision {
				digitsN := strings.Split(strconv.Itoa(n), "")
				digitsD := strings.Split(strconv.Itoa(d), "")
				if isInSlice(i, digitsN) && isInSlice(j, digitsD) && isDigitInCommon(digitsN, digitsD) {

					return true, i, j

				}
			}
		}
	}
	return false, 0, 0
}

func isInSlice(x int, s []string) bool {
	for _, v := range s {
		vInt, _ := strconv.Atoi(v)
		if x == vInt {
			return true
		}
	}
	return false
}

func isDigitInCommon(x, y []string) bool {
	for _, u := range x {
		for _, v := range y {
			if u == v {
				return true
			}
		}
	}
	return false
}
