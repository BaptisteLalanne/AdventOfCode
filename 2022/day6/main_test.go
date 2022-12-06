package main

import (
	"testing"
)

var example1 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
var example2 = `bvwbjplbgvbhsrlpgdmjqwftvncz`
var example3 = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  7,
		},
		{
			name:  "example2",
			input: example2,
			want:  5,
		},
		{
			name:  "example3",
			input: example3,
			want:  11,
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
			want:  19,
		},
		{
			name:  "example2",
			input: example2,
			want:  23,
		},
		{
			name:  "example3",
			input: example3,
			want:  26,
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
