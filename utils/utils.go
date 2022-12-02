package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadLines opens the input file and reads its lines.
func ReadLines(filename string) []string {
	f := mustOpen(filename)
	defer f.Close()
	s := bufio.NewScanner(f)
	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}
	if len(input) == 0 {
		Fatal("Empty input")
	}
	return input
}

// ReadInts opens the input file and reads its lines as ints.
func ReadInts(filename string) []int {
	f := mustOpen(filename)
	defer f.Close()
	s := bufio.NewScanner(f)
	var ints []int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			Fatal(err)
		}
		ints = append(ints, i)
	}
	if len(ints) == 0 {
		Fatal("Empty input")
	}
	return ints
}

func mustOpen(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		Fatal("Couldn't open file:", err)
	}
	return f
}

// Fatal exits the program after displaying the arguments on stderr
func Fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}
