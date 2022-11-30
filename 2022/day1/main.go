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
		utils.ToClipboard(fmt.Sprintf("%v", ans), "linux")
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		utils.ToClipboard(fmt.Sprintf("%v", ans), "linux")
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	nums := parseInput(input)
	fmt.Println(nums)
	return 0
}

func part2(input string) int {
	nums := parseInput(input)
	fmt.Println(nums)
	return 0
}

func parseInput(input string) (nums []int) {
	for _, v := range strings.Split(input, ",") {
		nums = append(nums, utils.ToInt(v))
	}
	return nums
}
