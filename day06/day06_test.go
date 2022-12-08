package day06

import "testing"

func Test_startOfPacket(t *testing.T) {
	type args struct {
		line   string
		msgLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"error 0", args{"", 4}, -1},
		{"error 1", args{"a", 4}, -1},
		{"error 2", args{"ab", 4}, -1},
		{"error 3", args{"abc", 4}, -1},
		{"error 4", args{"abcd", 4}, -1},
		{"test 1.1", args{"mjqjpqmgb", 4}, 7},
		{"test 1.2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4}, 5},
		{"test 1.3", args{"nppdvjthqldpwncqszvftbrmjlhg", 4}, 6},
		{"test 1.4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4}, 10},
		{"test 1.5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4}, 11},
		{"test 1.6", args{"abcde", 4}, 4},
		// Part 2
		{"test 2.1", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14}, 19},
		{"test 2.2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14}, 23},
		{"test 2.3", args{"nppdvjthqldpwncqszvftbrmjlhg", 14}, 23},
		{"test 2.4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14}, 29},
		{"test 2.5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfPacket(tt.args.line, tt.args.msgLen); got != tt.want {
				t.Errorf("startOfPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}
