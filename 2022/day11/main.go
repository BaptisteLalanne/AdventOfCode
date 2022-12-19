// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day11
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8          192326             30065 ns/op            7128 B/op        143 allocs/op
// BenchmarkPart2-8             302          19598393 ns/op            7384 B/op        144 allocs/op

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/BaptisteLalanne/AdventOfCode/utils"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		utils.ToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		utils.ToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

type monkey struct {
	items        []int
	operation    func(int) int
	divisor      int
	monkeyThrowT int
	monkeyThrowF int
}

const (
	rounds1 = 20
	rounds2 = 10000
)

func part1(input string) int {
	monkeys := parseInput(input)
	inspections := make([]int, len(monkeys))
	for r := 0; r < rounds1; r++ {
		for i := range monkeys {
			inspections[i] += len(monkeys[i].items)
			for _, item := range monkeys[i].items {
				item = monkeys[i].operation(item)
				item /= 3

				var nextMonkey int
				if item%monkeys[i].divisor == 0 {
					nextMonkey = monkeys[i].monkeyThrowT
				} else {
					nextMonkey = monkeys[i].monkeyThrowF
				}

				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, item)
			}
			monkeys[i].items = monkeys[i].items[:0]
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

func part2(input string) int {
	monkeys := parseInput(input)
	inspections := make([]int, len(monkeys))

	// Find the common divisor of all the monkeys
	products := 1
	for i := range monkeys {
		products *= monkeys[i].divisor
	}
	for r := 0; r < rounds2; r++ {
		for i := range monkeys {
			inspections[i] += len(monkeys[i].items)
			for _, item := range monkeys[i].items {
				item = monkeys[i].operation(item)
				// Reduice value of the item by keeping its divisable properties
				item %= products
				var nextMonkey int
				if item%monkeys[i].divisor == 0 {
					nextMonkey = monkeys[i].monkeyThrowT
				} else {
					nextMonkey = monkeys[i].monkeyThrowF
				}

				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, item)
			}
			monkeys[i].items = monkeys[i].items[:0]
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]

}

func parseInput(input string) (monkeys []monkey) {
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 7 {
		var items []int

		// Items
		strItems := strings.TrimPrefix(lines[i+1], "  Starting items: ")
		for _, str := range strings.Split(strItems, ", ") {
			item := utils.ToInt(str)
			items = append(items, item)
		}

		// Operator
		strFunction := strings.Split(strings.TrimPrefix(lines[i+2], "  Operation: new = old "), " ")
		var operation func(int) int
		switch strFunction[0] {
		case "+":
			switch strFunction[1] {
			case "old":
				operation = func(old int) int { return old + old }
			default:
				val := utils.ToInt(strFunction[1])
				operation = func(old int) int { return old + val }

			}
		case "*":
			switch strFunction[1] {
			case "old":
				operation = func(old int) int { return old * old }
			default:
				val := utils.ToInt(strFunction[1])
				operation = func(old int) int { return old * val }

			}
		}

		// Divisor
		strDivisor := strings.TrimPrefix(lines[i+3], "  Test: divisible by ")
		divisor := utils.ToInt(strDivisor)

		// Monkey to throw if True
		strMonkeyT := strings.TrimPrefix(lines[i+4], "    If true: throw to monkey ")
		monkeyThrowT := utils.ToInt(strMonkeyT)

		// Monkey to throw if False
		strMonkeyF := strings.TrimPrefix(lines[i+5], "    If false: throw to monkey ")
		monkeyThrowF := utils.ToInt(strMonkeyF)

		m := monkey{
			items:        items,
			operation:    operation,
			divisor:      divisor,
			monkeyThrowT: monkeyThrowT,
			monkeyThrowF: monkeyThrowF,
		}
		monkeys = append(monkeys, m)
	}
	return monkeys
}
