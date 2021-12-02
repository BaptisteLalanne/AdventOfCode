#!/usr/bin/env python3

"""Advent of Code 2021 Day 01"""


if __name__ == '__main__':
    with open("./inputs/day1.txt", 'r') as file:
        data = [int(line.strip()) for line in file.readlines()]

    increase = sum([1 for i in range(1, len(data)) if data[i] > data[i - 1]])
    print("Part 1",increase)
    increase = sum([1 for i in range(3, len(data)) if data[i] > data[i-3]])
    print("Part 2",increase)