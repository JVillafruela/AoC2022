package day01

import (
	"aoc2022/utils"
	"sort"
	"strconv"
)

func Part1(filename string) {
	lines := utils.ReadLines(filename)

	maxCalories := 0
	maxElve := 0
	curElve := 0
	sumCalories := 0
	linesCount := len(lines)

	for i, line := range lines {
		if len(line) == 0 || i == linesCount {
			curElve++
			if sumCalories > maxCalories {
				maxCalories = sumCalories
				maxElve = curElve
			}
			sumCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			utils.Fatal(err)
		}

		sumCalories += calories
	}
	println("Day 01 part1", maxElve, maxCalories)
}

func Part2(filename string) {
	lines := utils.ReadLines(filename)
	linesCount := len(lines)
	calories := make([]int, 0, linesCount) // indexed by elve#
	sumCalories := 0

	for i, line := range lines {

		cal, err := strconv.Atoi(line)
		blankLine := err != nil
		lastLine := i == linesCount-1

		if !blankLine {
			sumCalories += cal
		}

		if blankLine || lastLine {
			calories = append(calories, sumCalories)
			sumCalories = 0
			continue
		}

	}
	sort.Ints(calories)
	elvesCount := len(calories)
	for i := elvesCount - 3; i < elvesCount; i++ {
		sumCalories += calories[i]
	}
	println("Day 01 part2", sumCalories)
}
