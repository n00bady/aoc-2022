package main

import (
	"bufio"
	"fmt"
	"math"
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

func (tail *Point) follow(head Point) Point {
	x_diff := head.X - tail.X
	y_diff := head.Y - tail.Y

	if math.Abs(float64(x_diff)) <= 1 && math.Abs(float64(y_diff)) <= 1 {
		return *tail
	} else if math.Abs(float64(x_diff)) > 1 && math.Abs(float64(y_diff)) > 1 {
		return Point{tail.X + sign(x_diff), tail.Y + sign(y_diff)}
	} else if math.Abs(float64(x_diff)) > 1 {
		return Point{tail.X + sign(x_diff), head.Y}
	} else {
		return Point{head.X, tail.Y + sign(y_diff)}
	}
}

func sign(num int) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}

func solve() (int, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rope := []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	visitedA := make(map[Point]struct{})
	visitedA[rope[1]] = struct{}{}

	visitedB := make(map[Point]struct{})
	visitedB[rope[9]] = struct{}{}

	dirs := map[string]Point{"R": {1, 0}, "L": {-1, 0}, "U": {0, 1}, "D": {0, -1}}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		steps, _ := strconv.Atoi(line[1])

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].add(dirs[dir])
			for i, knot := range rope[1:] {
				rope[i+1] = knot.follow(rope[i])
			}

			visitedA[rope[1]] = struct{}{}
			visitedB[rope[9]] = struct{}{}
		}
	}

	return len(visitedA), len(visitedB)
}

func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %d\nB: %d\n", result_a, result_b)

}
