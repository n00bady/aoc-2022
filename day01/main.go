package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	topThree := [3]int{0, 0, 0}
	total := 0

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())

		if err != nil {
			if total > topThree[0] {
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = total
			} else if total > topThree[1] {
				topThree[2] = topThree[1]
				topThree[1] = total
			} else if total > topThree[2] {
				topThree[2] = total
			}

			total = 0
		} else {
			total += line
		}
	}

	topThreeSum := 0
	for _, element := range topThree {
		topThreeSum += element
	}

	return topThree[0], topThreeSum
}

func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %d\nB: %d\n", result_a, result_b)
}
