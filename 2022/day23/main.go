// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day23
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8             271          20751509 ns/op         1124849 B/op        674 allocs/op
// BenchmarkPart2-8               3        2107156333 ns/op        95537930 B/op      62284 allocs/op

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

type elf struct {
	choice coord
	move   bool
}

type coord struct {
	y int
	x int
}

func emptyGround(elves map[coord]elf, size [2]int) int {
	minX := size[1] - 1
	maxX := 0
	minY := size[0] - 1
	maxY := 0

	for k := range elves {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func printMap(elves map[coord]elf, size [2]int) {
	for y := 0; y < size[0]; y++ {
		for x := 0; x < size[1]; x++ {
			if _, ok := elves[coord{y, x}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

// Prefered direction
var pref = []string{"N", "S", "W", "E"}

func moveElves(elves map[coord]elf, rounds int) int {
	shift := 0
	moved := true

	for {
		if !moved && rounds == 0 || (rounds != 0 && shift == rounds) {
			break
		}

		moved = false

		for k, elf := range elves {
			// Reset choice
			elf.move = false

			x := k.x
			y := k.y

			canMove := false

			coordMap := map[string]coord{
				"N":  {y - 1, x},
				"NW": {y - 1, x - 1},
				"NE": {y - 1, x + 1},
				"S":  {y + 1, x},
				"SW": {y + 1, x - 1},
				"SE": {y + 1, x + 1},
				"W":  {y, x - 1},
				"E":  {y, x + 1},
			}

			for _, coords := range coordMap {
				if _, ok := elves[coords]; ok {
					canMove = true
					break
				}
			}

			// Select choice of movement
			if canMove {
				for i := 0; i < len(pref); i++ {
					var keys []string
					dir := pref[(i+shift)%len(pref)]
					switch dir {
					case "N":
						keys = []string{"N", "NW", "NE"}
					case "S":
						keys = []string{"S", "SW", "SE"}
					case "W":
						keys = []string{"W", "NW", "SW"}
					case "E":
						keys = []string{"E", "SE", "NE"}
					}

					valid := true
					for _, key := range keys {
						if _, around := elves[coordMap[key]]; around {
							valid = false
							break
						}
					}
					if valid {
						elf.choice = coordMap[dir]
						elf.move = true
						break
					}
				}
			}
			// Update the elf in the map
			elves[k] = elf
		}

		// Check if multiple elves want to move to the same cell
		cell := make(map[coord]int)
		for _, elf := range elves {
			if elf.move {
				cell[elf.choice]++
			}
		}

		// Update map
		for k, elf := range elves {
			if elf.move && cell[elf.choice] == 1 {
				// Choice can't be already taken (no need of a temp map)
				delete(elves, k)
				elves[elf.choice] = elf
				moved = true
			}
		}
		shift++
	}
	return shift
}

func part1(input string) int {
	elves, size := parseInput(input)
	moveElves(elves, 10)
	return emptyGround(elves, size)
}

func part2(input string) int {
	elves, _ := parseInput(input)
	return moveElves(elves, 0)
}

func parseInput(input string) (elves map[coord]elf, size [2]int) {
	lines := strings.Split(input, "\n")
	elves = make(map[coord]elf)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				newElf := elf{}
				elves[coord{y, x}] = newElf
			}
		}
	}

	size[0] = len(lines)
	size[1] = len(lines[0])

	return elves, size
}
