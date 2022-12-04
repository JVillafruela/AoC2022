package day04

import (
	"testing"
)

func Test_processLine1(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"included", args{"2-8,3-7"}, 1},
		{"not included", args{"2-4,6-8"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLine1(tt.args.line); got != tt.want {
				t.Errorf("processLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processLine2(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"included", args{"2-8,3-7"}, 0},
		{"overlap1", args{"2-4,4-8"}, 1},
		{"overlap2", args{"5-8,7-9"}, 1},
		{"disjointed", args{"1-3,4-6"}, 0},
		// In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap,
		// while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:
		{"N 2-4,6-8", args{"2-4,6-8"}, 0},
		{"N 2-3,4-5", args{"2-3,4-5"}, 0},
		{"Y 5-7,7-9", args{"5-7,7-9"}, 1},
		{"Y 2-8,3-7", args{"2-8,3-7"}, 1},
		{"Y 6-6,4-6", args{"6-6,4-6"}, 1},
		{"Y 2-6,4-8", args{"2-6,4-8"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLine2(tt.args.line); got != tt.want {
				t.Errorf("processLine2() = %v, want %v", got, tt.want)
			}
		})
	}
}
