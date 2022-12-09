// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day8
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           48826            107517 ns/op          106960 B/op        201 allocs/op
// BenchmarkPart2-8           10902            544896 ns/op          106960 B/op        201 allocs/op

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

func part1(input string) int {
	trees := parseInput(input)
	visible := make([][]bool, len(trees))
	for row := range trees {
		visible[row] = make([]bool, len(trees[row]))
	}

	// Find all visible trees from the TOP
	for col := 0; col < len(trees[0]); col++ {
		maxHeight := -1
		for row := 0; row < len(trees); row++ {
			height := trees[row][col]
			if height > maxHeight {
				visible[row][col] = true
				maxHeight = height
			}
		}
	}

	// Find all visible trees from the BOTTOM
	for col := 0; col < len(trees[0]); col++ {
		maxHeight := -1
		for row := len(trees) - 1; row >= 0; row-- {
			height := trees[row][col]
			if height > maxHeight {
				visible[row][col] = true
				maxHeight = height
			}
		}
	}

	// Find all visible trees from the LEFT
	for row := 0; row < len(trees); row++ {
		maxHeight := -1
		for col := 0; col < len(trees[0]); col++ {
			height := trees[row][col]
			if height > maxHeight {
				visible[row][col] = true
				maxHeight = height
			}
		}
	}

	// Find all visible trees from the RIGHT
	for row := 0; row < len(trees); row++ {
		maxHeight := -1
		for col := len(trees[0]) - 1; col >= 0; col-- {
			height := trees[row][col]
			if height > maxHeight {
				visible[row][col] = true
				maxHeight = height
			}
		}
	}

	count := 0
	for row := range visible {
		for col := range visible[row] {
			if visible[row][col] {
				count++
			}
		}
	}

	return count
}
func viewingDistance(trees *[][]int, row int, col int, direction string) int {
	var rowIncrement, colIncrement int

	switch direction {
	case "top":
		rowIncrement = 0
		colIncrement = -1
	case "bottom":
		rowIncrement = 0
		colIncrement = 1
	case "left":
		rowIncrement = -1
		colIncrement = 0
	case "right":
		rowIncrement = 1
		colIncrement = 0
	}

	count := 0
	maxHeight := (*trees)[row][col]
	row += rowIncrement
	col += colIncrement

	for (row < len(*trees) && row >= 0) && (col < len((*trees)[0]) && col >= 0) {
		height := (*trees)[row][col]
		// Count the tree even if it is higher
		count++

		if height >= maxHeight {
			break
		}

		row += rowIncrement
		col += colIncrement
	}
	return count
}

func scenicScore(trees *[][]int, row int, col int) int {
	return viewingDistance(trees, row, col, "top") *
		viewingDistance(trees, row, col, "bottom") *
		viewingDistance(trees, row, col, "left") *
		viewingDistance(trees, row, col, "right")
}

func part2(input string) int {
	trees := parseInput(input)
	visible := make([][]bool, len(trees))

	for row := range trees {
		visible[row] = make([]bool, len(trees[row]))
	}

	bestScenicScore := -1

	for row := range trees {
		for col := range trees[row] {
			score := scenicScore(&trees, row, col)
			if score > bestScenicScore {
				bestScenicScore = score
			}
		}
	}

	return bestScenicScore

}

func parseInput(input string) (tress [][]int) {
	lines := strings.Split(input, "\n")
	trees := make([][]int, len(lines))
	for i, line := range lines {
		trees[i] = make([]int, len(line))

		for j, char := range line {
			trees[i][j] = int(char - '0')
		}
	}

	return trees
}
