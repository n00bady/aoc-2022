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

	for scanner.Scan() {
		line := scanner.Text()

		var a1, a2, b1, b2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a1, &a2, &b1, &b2)

		if (a1 <= b1 && a2 >= b2) || (a1 >= b1 && a2 <= b2) {
			totalA++
			totalB++
		} else if (a1 >= b1 && a1 <= b2) || (b1 >= a1 && b1 <= a2) {
			totalB++
		}
	}

	return totalA, totalB
}

func main() {
	resultA, resultB := solve()
	fmt.Printf("A: %d\nB: %d\n", resultA, resultB)
}
