package main

import (
	"fmt"
	"slices"
	"strings"
)

var sp = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	heapPermutation(sp, len(sp))

	toSort := make([]string, 0)
	for i := 0; i < len(permutations); i++ {
		toSort = append(toSort, strings.Join(permutations[i], ""))
	}

	slices.Sort(toSort)
	fmt.Println(toSort[999999])
}

var permutations = make([][]string, 0)

func heapPermutation(a []string, size int) {
	if size == 1 {
		c := make([]string, len(a))
		copy(c, a)
		permutations = append(permutations, c)
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1)
		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
