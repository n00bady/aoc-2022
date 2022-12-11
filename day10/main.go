package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func tick(reg *int, cycle *int, sum *int, tmp *int, output *[6][40]string) {
	*reg += *tmp
	*tmp = 0

	trueCycle := *cycle - 1

	if math.Abs(float64(*reg-(trueCycle%40))) <= 1 {
		(*output)[trueCycle/40][trueCycle%40] = "#"
	} else {
		(*output)[trueCycle/40][trueCycle%40] = " "
	}

	if *cycle%40 == 20 {
		*sum += *reg * *cycle
	}

	*cycle++
}

func solve() (int, [6][40]string) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	sum, tmp := 0, 0
	cycle, reg := 1, 1
	var output [6][40]string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[0] == "noop" {
			tick(&reg, &cycle, &sum, &tmp, &output)
		} else if line[0] == "addx" {
			tick(&reg, &cycle, &sum, &tmp, &output)
			tick(&reg, &cycle, &sum, &tmp, &output)

			arg, _ := strconv.Atoi(line[1])
			tmp = arg
		}
	}

	return sum, output
}

func main() {
	resultA, resultB := solve()

	fmt.Printf("A: %d\nB:\n", resultA)
	for _, line := range resultB {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}
