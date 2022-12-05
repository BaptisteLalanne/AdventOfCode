// goos: linux
// goarch: amd64
// pkg: github.com/BaptisteLalanne/AdventOfCode/day5
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkPart1-8           36294            178338 ns/op          119512 B/op       2669 allocs/op
// BenchmarkPart2-8           22422            271266 ns/op          221033 B/op       4101 allocs/op

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

type Stack struct {
	elements []string
}

func (s *Stack) push(element string) {
	s.elements = append(s.elements, element)
}

func (s *Stack) pop() string {
	// Get the last element from the slice
	element := s.elements[len(s.elements)-1]
	// Remove the last element from the slice
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack) peek() string {
	return s.elements[len(s.elements)-1]
}

func move(stack1 *Stack, stack2 *Stack) {
	stack2.push(stack1.pop())
}

func moveKeepOrder(stack1 *Stack, stack2 *Stack, num int) {
	var tmpStack Stack
	for i := 0; i < num; i++ {
		tmpStack.push(stack1.pop())
	}
	for i := 0; i < num; i++ {
		stack2.push(tmpStack.pop())
	}
}

type Command struct {
	num         int
	source      int
	destination int
}

func part1(input string) string {
	stacks, commands := parseInput(input)
	for _, p := range commands {
		for j := 0; j < p.num; j++ {
			move(stacks[p.source], stacks[p.destination])
		}
	}

	var result string
	for _, stack := range stacks {
		result += stack.peek()
	}
	return result
}

func part2(input string) string {
	stacks, commands := parseInput(input)
	for _, p := range commands {
		moveKeepOrder(stacks[p.source], stacks[p.destination], p.num)
	}

	var result string
	for _, stack := range stacks {
		result += stack.peek()
	}
	return result
}

func parseInput(input string) (stacks []*Stack, commands []*Command) {
	data := strings.Split(input, "\n\n")

	// Handle stacks (data[0])
	lines := strings.Split(data[0], "\n")

	// Get last line to get the number of stacks
	lastLine := strings.Split(lines[len(lines)-1], " ")
	numberOfStacks := len(lastLine) / 3
	for i := 0; i < numberOfStacks; i++ {
		stacks = append(stacks, &Stack{})
	}

	lines = lines[:len(lines)-1]
	// Loop through the lines in descending order
	for i := len(lines) - 1; i >= 0; i-- {
		for pointer := 1; pointer < len(lines[i]); pointer += 4 {
			if lines[i][pointer] != ' ' {
				letter := utils.ToString(lines[i][pointer])
				// Get the number of the stack
				number := pointer / 4
				stacks[number].push(letter)
			}
		}
	}

	// Handle commands (data[1])
	lines = strings.Split(data[1], "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		command := Command{
			num:         utils.ToInt(words[1]),
			source:      utils.ToInt(words[3]) - 1,
			destination: utils.ToInt(words[5]) - 1,
		}
		commands = append(commands, &command)
	}
	return stacks, commands
}
