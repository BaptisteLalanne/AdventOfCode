// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day4
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           16489            364264 ns/op          176384 B/op       7001 allocs/op
// BenchmarkPart2-8           16711            367127 ns/op          176384 B/op       7001 allocs/op

package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) int {
	r := parseInput(input)

	nbOverlap := 0
	for _, r := range r {
		pair := strings.Split(r, ",")
		firstElf := strings.Split(pair[0], "-")
		secondElf := strings.Split(pair[1], "-")

		start1 := utils.ToInt(firstElf[0])
		end1 := utils.ToInt(firstElf[1])
		start2 := utils.ToInt(secondElf[0])
		end2 := utils.ToInt(secondElf[1])

		// First elf includes all second elf ids or the opposite
		if start1 >= start2 && end1 <= end2 || start2 >= start1 && end2 <= end1 {
			nbOverlap++
		}
	}
	return nbOverlap
}

func part2(input string) int {
	r := parseInput(input)

	nbOverlap := 0
	for _, r := range r {
		pair := strings.Split(r, ",")
		firstElf := strings.Split(pair[0], "-")
		secondElf := strings.Split(pair[1], "-")

		start1 := utils.ToInt(firstElf[0])
		end1 := utils.ToInt(firstElf[1])
		start2 := utils.ToInt(secondElf[0])
		end2 := utils.ToInt(secondElf[1])

		if start1 <= end2 && start2 <= start1 || start2 <= end1 && start1 <= start2 {
			nbOverlap++
		}
	}
	return nbOverlap
}

func parseInput(input string) (nums []string) {
	return strings.Split(input, "\n")
}
