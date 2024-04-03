package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var lettersValue = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func main() {
	names := getSortedNames()
	fmt.Println(getScores(names))
}

func getScores(names []string) int {
	sumScore := 0
	for i, n := range names {
		chars := strings.Split(n, "")
		score := 0
		for _, c := range chars {
			if c != "\"" {
				score += lettersValue[c]
			}
		}
		sumScore += score * (i + 1)
	}
	return sumScore
}

func getSortedNames() []string {
	// Open file.
	f, _ := os.Open("names.txt")
	defer f.Close()

	// Create a scanner to read from the file.
	scanner := bufio.NewScanner(f)

	// Read line by line.
	scanner.Scan()
	line := scanner.Text()
	names := strings.Split(line, ",")

	// Check for errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	slices.Sort(names)

	return names
}
