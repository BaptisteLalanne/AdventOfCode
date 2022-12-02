// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day2
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8            5120            228362 ns/op          120960 B/op       2501 allocs/op
// BenchmarkPart2-8            4621            266015 ns/op          120960 B/op       2501 allocs/op

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
	games := parseInput(input)
	score := 0
	for _, game := range games {
		move := strings.Split(game, " ")
		switch move[1] {
		case "X":
			score += 1 + playRPS(move[0], move[1])
		case "Y":
			score += 2 + playRPS(move[0], move[1])
		case "Z":
			score += 3 + playRPS(move[0], move[1])
		}
	}
	return score
}

var gameMoves = map[string]int{"A": 1, "B": 2, "C": 3}

func part2(input string) int {
	games := parseInput(input)
	score := 0
	for _, game := range games {
		move := strings.Split(game, " ")
		switch move[1] {
		case "X":
			score += (gameMoves[move[0]]+1)%3 + 1
		case "Y":
			score += 3
			score += gameMoves[move[0]]
		case "Z":
			score += 6
			score += (gameMoves[move[0]]+3)%3 + 1
		}
	}
	return score
}

func playRPS(player1 string, player2 string) int {
	// Draw as default
	output := 3
	switch player1 {
	case "A":
		if player2 == "Y" {
			output = 6
		} else if player2 == "Z" {
			output = 0
		}
	case "B":
		if player2 == "Z" {
			output = 6
		} else if player2 == "X" {
			output = 0
		}
	case "C":
		if player2 == "X" {
			output = 6
		} else if player2 == "Y" {
			output = 0
		}
	}
	return output
}

func parseInput(input string) (nums []string) {
	return strings.Split(input, "\n")
}
