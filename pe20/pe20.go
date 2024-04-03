package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	x := new(big.Int)
	x.SetString("1", 10)

	for i := 2; i <= 1000; i++ {
		x.Mul(x, big.NewInt(int64(i)))
	}

	digits := strings.Split(x.String(), "")
	digitsSum := 0

	for _, d := range digits {
		d_str, _ := strconv.Atoi(d)
		digitsSum += d_str
	}

	fmt.Println(x)
	fmt.Println(digitsSum)
}
