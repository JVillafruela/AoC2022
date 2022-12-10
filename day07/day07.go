package day07

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	empty = iota
	cd
	ls
	dir
	file
)

/*
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
*/

func Part1and2(filename string) {
	lines := utils.ReadLines(filename)
	total := 0
	state := empty
	root := NewDirectory(nil, "/")
	curDir := &root
	var name string

	dirCount := counDirectories(lines)
	dirs := make(map[*Directory]int, dirCount)

	for _, line := range lines {

		if line[0] == '$' {
			state, name = getCommand(line)
			if state == cd {
				curDir = curDir.changeDir(name)
				dirs[curDir] = 0
			}
			continue
		}

		if state == ls {
			name, size := getFile(line)
			if size == 0 {
				dir := NewDirectory(curDir, name)
				curDir.addChild(&dir)
			} else {
				curDir.addSize(size)
			}
		}
	}

	i := 0
	sizes := make([]int, len(dirs))
	for d := range dirs {
		dirs[d] = d.size()
		if dirs[d] < 100_000 {
			total += dirs[d]
		}
		sizes[i] = dirs[d]
		//fmt.Printf("Dir %s size %d\n", d.Name, sizes[i])
		i++
	}

	println("Day 07 part1", total)

	usedSpace := dirs[&root]
	freeSpace := 70_000_000 - usedSpace
	spaceToFree := 30_000_000 - freeSpace
	fmt.Printf("usedSpace %d spaceToFree %d\n", usedSpace, spaceToFree)
	// sort s
	sort.Ints(sizes)
	for _, size := range sizes {
		//fmt.Printf("size %d\n", size)
		if size > spaceToFree {
			println("Day 07 part2", size)
			break
		}
	}
}

func getCommand(line string) (state int, argument string) {
	// $ cd /  .. a
	if line[2:4] == "cd" {
		state = cd
		argument = line[5:]
	}
	if line[2:4] == "ls" {
		state = ls
	}
	return
}

func getFile(line string) (name string, size int) {
	if line[0:4] == "dir " {
		return line[4:], 0
	}
	if line[0:4] != "dir " {
		parts := strings.Split(line, " ")
		name := parts[1]
		size, _ := strconv.Atoi(parts[0])
		return name, size
	}
	return "", 0
}

func counDirectories(lines []string) int {
	count := 0
	for _, line := range lines {
		if line[0:3] == "dir" {
			count++
		}
	}
	return count

}
