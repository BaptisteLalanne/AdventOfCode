// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day3
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           14590            398651 ns/op          220710 B/op        799 allocs/op
// BenchmarkPart2-8           12666            472893 ns/op          291778 B/op        530 allocs/op

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

	score := 0
	for _, r := range r {
		first := strings.Split(r[:len(r)/2], "")
		second := strings.Split(r[len(r)/2:], "")
		charMap := make(map[string]bool)
		var commonChar string
		for _, char := range first {
			charMap[char] = true
		}
		for _, char := range second {
			if _, found := charMap[char]; found {
				commonChar = char
				break
			}
		}
		score += getPriorityValue(commonChar)
	}
	return score
}

func part2(input string) int {
	r := parseInput(input)
	score := 0
	for i := 0; i+2 < len(r); i += 3 {
		first := strings.Split(r[i], "")
		second := strings.Split(r[i+1], "")
		third := strings.Split(r[i+2], "")
		charMap := make(map[string]int)
		var commonChar string
		for _, char := range first {
			charMap[char] = 1
		}
		for _, char := range second {
			if _, found := charMap[char]; found {
				charMap[char] = 2
			}
		}
		for _, char := range third {
			if num := charMap[char]; num == 2 {
				commonChar = char
				break
			}
		}

		score += getPriorityValue(commonChar)
	}
	return score
}

func getPriorityValue(s string) int {
	ascii := utils.ToASCIICode(s)
	if ascii >= utils.ToASCIICode("a") {
		ascii = ascii - utils.ToASCIICode("a") + 1
	} else {
		ascii = ascii - utils.ToASCIICode("A") + 27
	}
	return ascii
}

func parseInput(input string) (nums []string) {
	return strings.Split(input, "\n")
}
