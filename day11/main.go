package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	op          func(int) int
	test        func(int) int
	inspections int
}

func parse() ([]Monkey, int) {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	monkeys := make([]Monkey, 0)
	modulo := 1

	// definitely not as neat as it could be, but gets the job done quite efficiently
	for _, s := range split {
		var items []int
		var opStr, itemsStr string
		var id, opNum, div, ifTrue, ifFalse int

		fixed := strings.Replace(strings.Replace(s, ", ", ",", -1), "old * old", "old ^ 2", 1)
		fmt.Sscanf(fixed, `Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`, &id, &itemsStr, &opStr, &opNum, &div, &ifTrue, &ifFalse)

		ops := map[string]func(int) int{
			"*": func(lvl int) int {
				return lvl * opNum
			},
			"^": func(lvl int) int {
				return int(math.Pow(float64(lvl), float64(opNum)))
			},
			"+": func(lvl int) int {
				return lvl + opNum
			},
		}

		test := func(lvl int) int {
			if lvl%div == 0 {
				return ifTrue
			} else {
				return ifFalse
			}
		}

		ss := strings.Split(itemsStr, ",")
		for _, s := range ss {
			num, _ := strconv.Atoi(s)
			items = append(items, num)
		}

		monkeys = append(monkeys, Monkey{items, ops[opStr], test, 0})
		modulo *= div
	}

	return monkeys, modulo
}

func solve(rounds int, noWorry bool) int {
	monkeys, modulo := parse()

	for i := 0; i < rounds; i++ {
		for j, m := range monkeys {
			for _, lvl := range m.items {
				lvl = m.op(lvl)

				if noWorry {
					lvl /= 3
				} else {
					lvl %= modulo
				}

				tgt := m.test(lvl)
				monkeys[tgt].items = append(monkeys[tgt].items, lvl)
			}

			monkeys[j].inspections += len(m.items)
			monkeys[j].items = nil
		}
	}

	sort.Slice(monkeys, func(i int, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}

func main() {
	resultA := solve(20, true)
	resultB := solve(10_000, false)
	fmt.Printf("A: %d\nB: %d\n", resultA, resultB)
}
