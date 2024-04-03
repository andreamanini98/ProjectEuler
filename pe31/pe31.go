package main

import "fmt"

func main() {
	var totWays int

	for a := 0; a < 2; a++ {
		for b := 0; b < 4; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 20; d++ {
					for e := 0; e < 40; e++ {
						for f := 0; f < 100; f++ {
							for g := 0; g < 200; g++ {
								if a*100+b*50+c*20+d*10+e*5+f*2+g == 200 {
									totWays++
								}
							}
						}
					}
				}
			}
		}
	}
	// The 7 here accounts for corner cases (all coins are the same)
	// The second 1 here accounts for the case of 200 itself.
	fmt.Println(totWays + 1 + 7)
}
