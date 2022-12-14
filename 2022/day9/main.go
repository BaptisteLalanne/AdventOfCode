// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day9
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8            4212           1277552 ns/op          630268 B/op       4156 allocs/op
// BenchmarkPart2-8            4206           1366653 ns/op          444930 B/op       4091 allocs/op

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
		ans := part2(input)
		utils.ToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

type move struct {
	dir   string
	count int
}
type coord struct {
	x int
	y int
}

func needToMove(head, tail *coord) bool {
	xDist := math.Abs(float64(head.x - tail.x))
	yDist := math.Abs(float64(head.y - tail.y))
	return xDist > 1 || yDist > 1
}

func moveTail(head, tail *coord) bool {
	var moved bool
	if needToMove(head, tail) {
		xDist := head.x - tail.x
		yDist := head.y - tail.y
		if xDist > 0 {
			tail.x++
		} else if xDist < 0 {
			tail.x--
		}
		if yDist > 0 {
			tail.y++
		} else if yDist < 0 {
			tail.y--
		}
		moved = true
	} else {
		moved = false
	}
	return moved
}

func part1(input string) int {
	moves := parseInput(input)

	seen := make(map[coord]bool)
	var head, tail coord

	count := 0

	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			switch move.dir {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}
			moveTail(&head, &tail)
			if _, ok := seen[tail]; !ok {
				count++
				seen[tail] = true
			}
		}
	}

	return count
}

func part2(input string) int {
	moves := parseInput(input)

	const (
		numberNodes = 10
	)
	seen := make(map[coord]bool)
	nodes := make([]coord, numberNodes)

	// Pointers on head & tail
	lastNode := &nodes[numberNodes-1]
	head := &nodes[0]

	count := 0

	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			switch move.dir {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}
			temp := head
			// Loop trough all nodes
			// If one node doesn't move, all nodes after won't
			for i, _ := range nodes {
				node := &nodes[i]
				if moved := moveTail(temp, node); !moved && i != 0 {
					break
				}
				temp = node
			}
			if _, ok := seen[*lastNode]; !ok {
				count++
				seen[*lastNode] = true
			}

		}
	}
	return count

}

func parseInput(input string) (moves []move) {
	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")
		move := move{words[0], utils.ToInt(words[1])}
		moves = append(moves, move)
	}
	return moves
}
