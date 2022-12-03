package day02

import (
	"aoc2022/utils"
	"fmt"
)

func Part1(filename string) {
	lines := utils.ReadLines(filename)

	score := 0
	var a, b rune

	for _, line := range lines {
		fmt.Sscanf(line, "%c %c", &a, &b)
		score += outcome1(a, b)
	}
	println("Day 02 part1", score)
}

// Elf   : A for Rock, B for Paper, and C for Scissors.
// Me    : X for Rock, Y for Paper, and Z for Scissors.
// score : score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func outcome1(a, b rune) int {
	const (
		rock     = 1
		paper    = 2
		scissors = 3
	)
	shapes := map[rune]int{'A': rock, 'B': paper, 'C': scissors, 'X': rock, 'Y': paper, 'Z': scissors}

	x := shapes[a] // elf
	y := shapes[b] // me
	if x == 0 || y == 0 {
		utils.Fatal("bad shape")
	}

	if x == y {
		return 3 + y
	}

	won := (x == rock && y == paper) || (x == paper && y == scissors) || (x == scissors && y == rock)

	if won {
		return 6 + y
	}

	return y
}

func Part2(filename string) {
	lines := utils.ReadLines(filename)

	score := 0
	var a, b rune

	for _, line := range lines {
		fmt.Sscanf(line, "%c %c", &a, &b)
		score += outcome2(a, b)
	}
	println("Day 02 part2", score)
}

// Elf   : A for Rock, B for Paper, and C for Scissors.
// Me    : X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
func outcome2(a, b rune) int {
	const (
		rock     = 'A'
		paper    = 'B'
		scissors = 'C'
	)

	const (
		lose = 'X'
		draw = 'Y'
		win  = 'Z'
	)

	winShapes := map[rune]rune{rock: paper, paper: scissors, scissors: rock}
	loseShapes := map[rune]rune{rock: scissors, paper: rock, scissors: paper}

	var shape rune
	switch b {
	case draw:
		shape = a
	case lose:
		shape = loseShapes[a]
	case win:
		shape = winShapes[a]
	}

	return outcome1(a, shape)
}
