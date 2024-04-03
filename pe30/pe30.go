package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for i := 0; i < 100000000; i++ {
		if isFifthSum(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}

func isFifthSum(x int) bool {
	xStr := strconv.Itoa(x)
	xDigits := strings.Split(xStr, "")

	sum := 0
	for _, d := range xDigits {
		dInt, _ := strconv.Atoi(d)
		sum += int(math.Pow(float64(dInt), 5))
	}
	return sum == x
}
