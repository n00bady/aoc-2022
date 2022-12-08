package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func appendParentSums(dirSums *map[string]int, path []string, filesize int) {
	for len(path) > 0 {
		(*dirSums)[strings.Join(path, "/")] += filesize
		path = path[:len(path)-1]
	}
}

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var path []string
	dirSums := make(map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[1] == "cd" {
			switch line[2] {
			case "..":
				path = path[:len(path)-1]
			case "/":
				path = []string{"/"}
			default:
				path = append(path, line[2])
				curStr := strings.Join(path, "/")

				if _, ok := dirSums[curStr]; !ok {
					dirSums[curStr] = 0
				}
			}
		} else if line[1] == "ls" {
			continue
		} else {
			filesize, _ := strconv.Atoi(line[0])
			tmp := make([]string, len(path))
			copy(tmp, path)

			appendParentSums(&dirSums, tmp, filesize)
		}
	}

	const limitA = 100000
	limitB := 30_000_000 - (70_000_000 - dirSums["/"])
	smallest := 70_000_000
	totalOver := 0

	for _, size := range dirSums {
		if size <= limitA {
			totalOver += size
		}

		if size >= limitB && size < smallest {
			smallest = size
		}
	}

	return totalOver, smallest
}

func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %d\nB: %d\n", result_a, result_b)
}
