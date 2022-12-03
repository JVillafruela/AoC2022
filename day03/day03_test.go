package day03

import (
	"testing"
)

func Test_index(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{'a'}, 0},
		{"b", args{'b'}, 1},
		{"z", args{'z'}, 25},
		{"A", args{'A'}, 26},
		{"B", args{'B'}, 27},
		{"Z", args{'Z'}, 26*2 - 1},
		{"é", args{'é'}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := index(tt.args.c); got != tt.want {
				t.Errorf("index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_positionASCII(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{0}, 'a'},
		{"b", args{1}, 'b'},
		{"z", args{25}, 'z'},
		{"A", args{26}, 'A'},
		{"B", args{27}, 'B'},
		{"Z", args{51}, 'Z'},
		{"ko 52", args{52}, -1},
		{"ko -1", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := positionASCII(tt.args.i); got != tt.want {
				t.Errorf("positionASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
