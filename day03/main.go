package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalA := 0
	totalB := 0
	counter := 1
	var topThree []string

	for scanner.Scan() {
		line := scanner.Text()
		topThree = append(topThree, line)
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		if counter%3 == 0 && counter != 0 {
			commonB := make(map[rune]bool)

			for _, char := range topThree[0] {
				if strings.Contains(topThree[1], string(char)) && strings.Contains(topThree[2], string(char)) && !commonB[char] {
					commonB[char] = true

					if int(char)-64 > 26 {
						totalB += int(char) - 64 - 32
					} else {
						totalB += int(char) - 64 + 26
					}

					break
				}
			}

			topThree = topThree[:0]
		}

		commonA := make(map[rune]bool)

		for _, char := range compartment1 {
			if strings.Contains(compartment2, string(char)) && !commonA[char] {
				commonA[char] = true

				if int(char)-64 > 26 {
					totalA += int(char) - 64 - 32
				} else {
					totalA += int(char) - 64 + 26
				}
			}
		}

		counter++
	}

	return totalA, totalB
}

func main() {
	resultA, resultB := solve()
	fmt.Printf("A: %d\nB: %d\n", resultA, resultB)
}
