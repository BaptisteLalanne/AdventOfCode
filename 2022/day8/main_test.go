package main

import (
	"testing"
)

var example1 = `30373
25512
65332
33549
35390`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  21,
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
			name:  "example1",
			input: example1,
			want:  8,
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
		r = part1(input)
	}
	result = r
}

func BenchmarkPart2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = part2(input)
	}
	result = r
}
