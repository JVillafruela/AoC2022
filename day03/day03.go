package day03

import (
	"aoc2022/utils"
)

const (
	a int = int('a')
	z int = int('z')
	A int = int('A')
	Z int = int('Z')
)

type backpack [52]int

func Part1(filename string) {
	lines := utils.ReadLines(filename)
	total := 0

	for _, line := range lines {
		total += processLine(line)
	}
	println("Day 03 part1", total)
}

func Part2(filename string) {
	var bp [3]backpack
	lines := utils.ReadLines(filename)
	total := 0

	for l, line := range lines {
		j := l % 3
		bp[j] = countItems(line)
		if j == 2 {
			for i, k := range bp[0] {
				if (k == 1) && (bp[1][i] == 1) && (bp[2][i] == 1) {
					priority := i + 1
					total += priority
					//fmt.Printf("priority %d %d %c\n", i, priority, positionASCII(i))
					break
				}
			}

		}
	}
	println("Day 03 part2", total)
}

func processLine(line string) int {
	k := len(line) / 2
	part1 := line[0:k]
	part2 := line[k:]

	items1 := countItems(part1)
	items2 := countItems(part2)

	for i, k := range items1 {
		if (k == 1) && (items2[i] == 1) {
			priority := i + 1
			//fmt.Printf("priority %d %d %c\n", i, priority, positionASCII(i))
			return priority
		}
	}

	return 0
}

func countItems(str string) backpack {
	var bp backpack
	for _, c := range str {
		i := index(c)
		bp[i] = 1
	}
	return bp
}

// compute position of character c in backpack
func index(c rune) int {

	x := int(c) // rune to int
	y := -1

	if a <= x && x <= z {
		y = x - a
	}
	if A <= x && x <= Z {
		y = (z - a + 1) + (x - A)
	}

	return y
}

// compute ASCII code from position in backpack
func positionASCII(i int) int {
	if 0 <= i && i <= z-a {
		return i + a
	}
	if z-a+1 <= i && i <= (z-a+1)+(Z-A) {
		return A + i - (z - a + 1)
	}

	return -1
}
