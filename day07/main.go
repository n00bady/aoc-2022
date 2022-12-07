package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var cur []string
	dirSums := make(map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if filesize, err := strconv.Atoi(line[0]); err == nil {
			tmp := make([]string, len(cur))
			copy(tmp, cur)

			// add sum to all parent directories
			for len(tmp) > 0 {
				dirSums[strings.Join(tmp, "/")] += filesize
				tmp = tmp[:len(tmp)-1]
			}
		} else if line[1] == "cd" {
			if line[2] == ".." && len(cur) > 1 {
				cur = cur[:len(cur)-1]
			} else if line[2] == "/" {
				cur = []string{"/"}
			} else {
				cur = append(cur, line[2])
				curStr := strings.Join(cur, "/")

				if _, ok := dirSums[curStr]; !ok {
					dirSums[curStr] = 0
				}
			}
		}
	}

	const limitA = 100000
	limitB := 30_000_000 - (70_000_000 - dirSums["/"])
	smallest := 70_000_000
	total := 0

	for _, size := range dirSums {
		if size <= limitA {
			total += size
		}

		if size >= limitB && size < smallest {
			smallest = size
		}
	}

	return total, smallest
}

func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %d\nB: %d\n", result_a, result_b)
}
