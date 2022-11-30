package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

/*
* CREDIT: https://github.com/alexchao26/advent-of-code-go
* Thanks for Alex Chao to share his utils AoC functions
* I just tweaked a little bit some of them.
**/
func ToClipboard(text string, arch string) error {
	var command *exec.Cmd
	// Mac "OS"
	if arch == "darwin" {
		command = exec.Command("pbcopy")
	}
	// Linux
	if arch == "linux" {
		command = exec.Command("xclip", "-selection", "c")
	}
	command.Stdin = bytes.NewReader([]byte(text))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting copy command: %w", err)
	}

	err := command.Wait()
	if err != nil {
		return fmt.Errorf("error running copy %w", err)
	}

	return nil
}

// ToInt will case a given arg into an int type.
// Supported types are:
//    - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToString will case a given arg into an int type.
// Supported types are:
//    - int
//    - byte
//    - rune
func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}
