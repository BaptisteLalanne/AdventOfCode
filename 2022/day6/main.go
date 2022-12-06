// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day6
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           20230            310338 ns/op          147680 B/op       3692 allocs/op
// BenchmarkPart2-8            9450            579795 ns/op          230556 B/op       5714 allocs/op

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

func findIndexDifferentsChar(input string, numDiff int) int {
	m := make(map[byte]bool)
	for i := 0; i < len(input)+1; i++ {
		m[input[i]] = true
		for j := i + 1; j < len(input)+1; j++ {
			if m[input[j]] {
				break
			} else {
				m[input[j]] = true
			}
			if len(m) == numDiff {
				return j + 1
			}
		}
		m = make(map[byte]bool)
	}
	return 0
}

func part1(input string) int {
	return findIndexDifferentsChar(input, 4)
}

func part2(input string) int {
	return findIndexDifferentsChar(input, 14)
}
