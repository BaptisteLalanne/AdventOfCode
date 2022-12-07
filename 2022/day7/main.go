// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day7
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           19423            296335 ns/op          119194 B/op       2260 allocs/op
// BenchmarkPart2-8           19402            300030 ns/op          119202 B/op       2260 allocs/op

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

func exploreFs(input string) map[string]int {
	input = "\n" + input
	r := parseInput(input)

	// Keys will be full path (to be unique)
	foldersSize := make(map[string]int)
	parentFolder := make(map[string]string)

	currentFolder := "/"

	for _, output := range r {
		lines := strings.Split(output, "\n")
		cmd := strings.Split(lines[0], " ")
		lines = lines[1:]
		switch cmd[0] {
		case "cd":
			switch cmd[1] {
			case "/":
				currentFolder = "/"
			case "..":
				currentFolder = parentFolder[currentFolder]
			default:
				// Construct the full path of the folder by appending
				// the folder name to the current folder path
				currentFolder = currentFolder + "/" + cmd[1]
			}

		case "ls":
			for _, line := range lines {
				words := strings.Split(line, " ")
				if words[0] != "dir" {
					size := utils.ToInt(words[0])
					temp := currentFolder

					// Loop to add file's size to all parent folders
					for exist := true; exist; temp, exist = parentFolder[temp] {
						foldersSize[temp] += size
					}
				} else {
					folderPath := currentFolder + "/" + words[1]
					parentFolder[folderPath] = currentFolder
				}
			}
		}

	}
	return foldersSize
}

func part1(input string) int {
	const maxSize = 100000

	foldersSize := exploreFs(input)

	total := 0
	for _, v := range foldersSize {
		if v <= maxSize {
			total += v
		}
	}
	return total
}

func part2(input string) int {
	const (
		fileSystemSize = 70000000
		updateSize     = 30000000
	)

	foldersSize := exploreFs(input)

	neededSize := updateSize - (fileSystemSize - foldersSize["/"])
	minFolderSize := foldersSize["/"]
	for _, v := range foldersSize {
		if v <= minFolderSize && v >= neededSize {
			minFolderSize = v
		}
	}
	return minFolderSize
}

func parseInput(input string) (cmds []string) {
	return strings.Split(input, "\n$ ")
}
