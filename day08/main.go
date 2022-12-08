package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p1 *Point) add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func getForest() map[Point]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	trees := make(map[Point]int)
	row := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for i, height := range line {
			trees[Point{i, row}], _ = strconv.Atoi(height)
		}

		row++
	}

	return trees
}

func solve() (int, int) {
	trees := getForest()
	dirs := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	top := -1
	vis := make(map[Point]struct{})

	for point, src := range trees {
		score := 1

		for _, dir := range dirs {
			tmp := dir

			for i := 1; ; i++ {
				cmp, ok := trees[point.add(tmp)]

				if !ok { // visible
					if i == 1 {
						score = 0
					} else {
						score *= i - 1
					}

					vis[point] = struct{}{}
					break
				} else if cmp >= src { // not visible
					score *= i
					break
				}

				tmp = tmp.add(dir)
			}
		}

		if score > top {
			top = score
		}
	}

	return len(vis), top
}

func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %d\nB: %d\n", result_a, result_b)
}
