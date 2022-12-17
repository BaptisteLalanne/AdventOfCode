// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day10
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8          273529             20277 ns/op           11197 B/op        113 allocs/op
// BenchmarkPart2-8          280048             22079 ns/op           11196 B/op        113 allocs/op

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
		part2(input)
	}
}

func part1(input string) int {
	instructions := parseInput(input)

	cycle := 0
	updates := make(map[int]int)
	for _, inst := range instructions {
		switch {
		case inst == "noop":
			cycle++
		case strings.HasPrefix(inst, "addx "):
			num := utils.ToInt(strings.TrimPrefix(inst, "addx "))
			cycle += 2
			updates[cycle] = num
		}
	}

	// Get values for each cycle
	currentVal := 1
	values := make([]int, cycle)
	for i := 0; i < cycle; i++ {
		if _, ok := updates[i]; ok {
			currentVal += updates[i]
		}
		values[i] = currentVal
	}

	sum := 0
	for c := 20; c <= 240; c += 40 {
		sum += c * values[c-1]
	}
	return sum
}

const (
	cols   = 40
	rows   = 6
	pixels = cols * rows
)

func part2(input string) {
	instructions := parseInput(input)

	cycle := 0
	updates := make(map[int]int)
	for _, inst := range instructions {
		switch {
		case inst == "noop":
			cycle++
		case strings.HasPrefix(inst, "addx "):
			num := utils.ToInt(strings.TrimPrefix(inst, "addx "))
			cycle += 2
			updates[cycle] = num
		}
	}

	// Get values for each cycle
	currentVal := 1
	values := make([]int, cycle)
	for i := 0; i < cycle; i++ {
		if _, ok := updates[i]; ok {
			currentVal += updates[i]
		}
		values[i] = currentVal
	}

	// Print the image
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			value := values[row*cols+col]
			if math.Abs(float64(value-col)) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

}

func parseInput(input string) (instr []string) {
	return strings.Split(input, "\n")
}
