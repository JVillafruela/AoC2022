package day10

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

type Register struct {
	X     int
	Start int
	End   int
}

func Part1(filename string) {
	lines := utils.ReadLines(filename)
	total := 0
	count := len(lines)
	run := make([]Register, count+1)
	x := 1
	var time, start, end, val int
	for i, line := range lines {
		op := line[0:4] // noop addx
		switch op {
		case "noop":
			time = 1
			val = 0
		case "addx":
			time = 2
			val, _ = strconv.Atoi(line[5:])
		}
		start = end + 1
		end = start + time - 1

		run[i] = Register{X: x, Start: start, End: end}
		x += val
	}
	run[count] = Register{X: x, Start: end + 1, End: end + 1}

	for i, r := range run {
		fmt.Printf("%d X=%d [%d,%d] \n", i, r.X, r.Start, r.End)
	}

	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		signal := signalStrength(run, i)
		total += signal
		fmt.Printf("i=%d, signal=%d\n", i, signal)
	}

	println("Day 10 part1", total)
}

func signalStrength(run []Register, cycle int) int {
	for _, r := range run {
		if r.Start <= cycle && cycle <= r.End {
			return r.X * cycle
		}
	}
	return -1
}
