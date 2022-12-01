package main

import (
	"testing"
)

var example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func BenchmarkPart1(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = part1(input)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkPart2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = part2(input)
	}
	result = r
}
