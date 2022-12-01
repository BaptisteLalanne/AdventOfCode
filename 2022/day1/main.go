// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day1
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkParsing-8        191590             30519 ns/op
// BenchmarkPart1-8        1000000000           1.781 ns/op
// BenchmarkPart2-8        215486608            25.81 ns/op

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
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

var parsedInput []string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	parsedInput = parseInput(input)
	if part == 1 {
		ans := part1(parsedInput)
		utils.ToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(parsedInput)
		utils.ToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(elfs []string) int {
	maximum := 0
	for _, elf := range elfs {
		foods := strings.Split(elf, "\n")
		total := 0
		for _, cal := range foods {
			total += utils.ToInt(cal)
		}
		if total > maximum {
			maximum = total
		}
	}
	return maximum
}

func part2(elfs []string) int {
	highest := make([]int, 3)
	for _, elf := range elfs {
		foods := strings.Split(elf, "\n")
		total := 0
		for _, cal := range foods {
			total += utils.ToInt(cal)
		}
		if total > highest[0] {
			// Get rid of the lowest then sort to find the new lowest
			highest[0] = total
			sort.Ints(highest)
		}
	}
	total := 0
	for _, highCal := range highest {
		total += highCal
	}
	return total
}

func parseInput(input string) (nums []string) {
	return strings.Split(input, "\n\n")
}
