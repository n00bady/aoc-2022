package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(markerLen int) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	char := strings.Split(scanner.Text(), "")
	set := make(map[string]struct{})
	var cur []string
	i := 0

	for len(set) < markerLen {
		set = make(map[string]struct{})
		cur, char = char[:markerLen], char[1:]

		for _, char := range cur {
			set[char] = struct{}{}
		}

		i++
	}

	return i + markerLen - 1
}

func main() {
	fmt.Printf("A: %d\nB: %d\n", solve(4), solve(14))
}
