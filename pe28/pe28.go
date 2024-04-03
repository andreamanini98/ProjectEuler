package main

import "fmt"

func main() {
	fmt.Println(getSpiralSum())
}

func getSpiralSum() int {
	// sum:     1 + 2 + 2 + 2 + 2 + 4 + 4 + 4 + 4 + 6 + 6 + 6 + 6
	// size:    1   3   3   3   3   5   5   5   5   7   7   7   7
	sum := 1
	n := 1

	for size := 3; size <= 1001; size += 2 {
		for i := 0; i < 4; i++ {
			n += (size - 1)
			sum += n
		}
	}
	return sum
}
