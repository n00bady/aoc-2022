package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p1 *Point) add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point) eq(p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (p *Point) valid(maxX int, maxY int) bool {
	return p.X >= 0 && p.X <= maxX && p.Y >= 0 && p.Y <= maxY
}

func bfs(grid []string, points []Point, goal Point) int {
	adjacents := []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	totalShortest := 10_000_000

	var cur, next Point
	var ch rune

	for _, point := range points {
		curShortest := 10_000_000

		dist := make(map[Point]int)
		queue := []Point{point}

		for len(queue) > 0 {
			cur = queue[0]
			queue = queue[1:]

			ch = rune(grid[cur.Y][cur.X])

			if cur.eq(goal) && dist[cur] < curShortest {
				curShortest = dist[cur]
			}

			for _, adj := range adjacents {
				next = cur.add(adj)
				_, exists := dist[next]

				if !next.valid(len(grid[0])-1, len(grid)-1) {
					continue
				} else if int(grid[next.Y][next.X])-int(ch) <= 1 && !exists {
					queue = append(queue, next)
					dist[next] = dist[cur] + 1
				}
			}
		}

		if curShortest < totalShortest {
			totalShortest = curShortest
		}
	}

	return totalShortest
}

func solve() (int, int) {
	input, _ := os.ReadFile("input.txt")
	grid := strings.Split(strings.TrimSpace(string(input)), "\n")

	starts := make([]Point, 0)
	var start, goal Point

	for y, row := range grid {
		for x, char := range row {
			if string(char) == "S" {
				start.X, start.Y = x, y
				starts = append(starts, Point{x, y})
				grid[y] = strings.Replace(grid[y], "S", "a", 1)
			} else if string(char) == "E" {
				goal.X, goal.Y = x, y
				grid[y] = strings.Replace(grid[y], "E", "z", 1)
			}

			if string(char) == "a" {
				starts = append(starts, Point{x, y})
			}
		}
	}

	resultA := bfs(grid, []Point{start}, goal)
	resultB := bfs(grid, starts, goal)

	return resultA, resultB
}

func main() {
	resultA, resultB := solve()
	fmt.Printf("A: %d\nB: %d\n", resultA, resultB)
}
