package day05

import (
	"aoc2022/utils"
	"fmt"
)

func Part1(filename string) {
	msg := process(filename, 9000)
	println("Day 04 part12", msg)
}

func Part2(filename string) {
	msg := process(filename, 9001)
	println("Day 04 part2", msg)
}

func process(filename string, craneModel int) string {
	lines := utils.ReadLines(filename)

	stackCount, stackCapacity, startMoves := getLimits(lines)
	//fmt.Printf("stackCount %d, stackCapacity %d, startMoves %d \n", stackCount, stackCapacity, startMoves)

	stacks := initStacks(stackCount, stackCapacity)

	if err := makeStacks(&stacks, lines, startMoves); err != nil {
		utils.Fatal(err)
	}
	if err := moveStacks(&stacks, lines, startMoves, craneModel); err != nil {
		fmt.Printf("%v\n", stacks)
		utils.Fatal(err)
	}

	return message(&stacks)
}

/*
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
*/
func getLimits(lines []string) (stackCount int, stackCapacity int, startMoves int) {
	for i, line := range lines {
		//fmt.Printf("DDD %d '%s' '%c' len %d \n", i, line, line[1], len(line))
		if line[1] == '1' {
			stackCount = (len(lines[i-1]) + 1) / 4
			stackCapacity = i
			startMoves = i + 2
			return
		}
	}
	return -1, -1, -1
}

func initStacks(stackCount int, stackCapacity int) []RuneStack {
	stacks := make([]RuneStack, 0, stackCount)
	for i := 0; i < stackCount; i++ {
		s := NewRuneStack(stackCapacity)
		stacks = append(stacks, s)
	}
	return stacks
}

func makeStacks(stacks *[]RuneStack, lines []string, startMoves int) error {

	lastStackLine := startMoves - 3
	for i := 0; i < startMoves-2; i++ {
		s := 0
		k := lastStackLine - i
		for j := 1; j < len(lines[i]); j += 4 {
			var r rune
			if _, err := fmt.Sscanf(lines[i][j:j+1], "%c", &r); err != nil {
				return fmt.Errorf("error scanning line %d position %d", i, j)
			}
			//fmt.Printf("Line %d stack %d index %d char %c\n", i, s, k, r)
			if 'A' <= r && r <= 'Z' {
				(*stacks)[s].poke(k, r)
			}
			s++
		}
		//println()
	}
	return nil
}

func moveStacks(stacks *[]RuneStack, lines []string, startMoves int, craneModel int) error {

	for i := startMoves; i < len(lines); i++ {
		//fmt.Printf("stacks %v\n", stacks)
		//fmt.Printf("line %d %s\n", i+1, lines[i])
		//move 1 from 2 to 1
		var count, from, to int
		n, err := fmt.Sscanf(lines[i], "move %d from %d to %d", &count, &from, &to)
		if err != nil || n != 3 {
			return fmt.Errorf("error scanning line %d n=%d %v", i, n, err)
		}
		//fmt.Printf("%s : %d %d %d\n", lines[i], count, from, to)
		crates, err := (*stacks)[from-1].popMany(count)
		if err != nil {
			return fmt.Errorf("error line %d '%s' %v ", i+1, lines[i], err)
		}
		if craneModel == 9001 {
			reverse(crates)
		}

		(*stacks)[to-1].pushMany(crates)
	}

	return nil
}

func reverse(a []rune) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func message(stacks *[]RuneStack) string {
	var runes []rune
	for _, s := range *stacks {
		runes = append(runes, s.top())
	}
	return string(runes)
}
