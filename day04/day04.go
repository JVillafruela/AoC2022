package day04

import (
	"aoc2022/utils"
	"fmt"
)

func Part1(filename string) {
	lines := utils.ReadLines(filename)
	total := 0

	for _, line := range lines {
		total += processLine1(line)
	}
	println("Day 04 part1", total)
}

func Part2(filename string) {
	lines := utils.ReadLines(filename)
	total := 0

	for _, line := range lines {
		total += processLine2(line)
	}
	println("Day 04 part2", total)
}

func scan(line string) (int, int, int, int) {
	var b1, e1, b2, e2 int
	if _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &b1, &e1, &b2, &e2); err != nil {
		utils.Fatal("Error scanning line " + line)
	}
	return b1, e1, b2, e2
}

// check for inclusion
func processLine1(line string) int {
	b1, e1, b2, e2 := scan(line)

	if (b1 <= b2 && b2 <= e2 && e2 <= e1) || (b2 <= b1 && b1 <= e1 && e1 <= e2) {
		return 1
	}

	return 0
}

// check for overlap
func processLine2(line string) int {
	b1, e1, b2, e2 := scan(line)

	if processLine1(line) == 1 {
		return 1
	}

	if (b2 <= e1 && e1 < e2) || (b2 < b1 && b1 <= e2) {
		return 1
	}

	return 0
}
