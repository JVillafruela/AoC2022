package day06

import (
	"aoc2022/utils"
)

func Part1(filename string) {
	n := process(filename, 4)
	println("Day 06 part1", n)
}

func Part2(filename string) {
	n := process(filename, 14)
	println("Day 06 part2", n)
}

func process(filename string, msgLen int) int {
	lines := utils.ReadLines(filename)
	return startOfPacket(lines[0], msgLen)

}

// identify the first position where the four most recently received characters were all different.
// Specifically, it needs to report the number of characters from the beginning of the buffer
// to the end of the first such four-character marker.
// 0123456789
// mjqjpqmgb
//    ^^^^ "jpqm"
func startOfPacket(line string, msgLen int) int {
	l := len(line)
	if l <= msgLen {
		return -1
	}
mainloop:
	for i := msgLen - 1; i < l; i++ {
		message := line[i-msgLen+1 : i+1]
		chars := make(map[rune]int, msgLen)
		for _, r := range message {
			chars[r]++
			if chars[r] == 2 {
				continue mainloop
			}
		}
		return i + 1
	}
	return -1
}
