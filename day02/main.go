package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalA := 0
	totalB := 0

	combinations := map[string][2]int{
		"A X": {4, 3},
		"B X": {1, 1},
		"C X": {7, 2},
		"A Y": {8, 4},
		"B Y": {5, 5},
		"C Y": {2, 6},
		"A Z": {3, 8},
		"B Z": {9, 9},
		"C Z": {6, 7},
	}

	for scanner.Scan() {
		line := scanner.Text()

		totalA += combinations[line][0]
		totalB += combinations[line][1]
	}

	return totalA, totalB
}

func main() {
	resultA, resultB := solve()
	fmt.Printf("A: %d\nB: %d\n", resultA, resultB)
}
