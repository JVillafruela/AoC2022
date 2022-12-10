package day07

import (
	"testing"
)

func Test_getCommand(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name         string
		args         args
		wantState    int
		wantArgument string
	}{
		{"cd /", args{"$ cd /"}, cd, "/"},
		{"cd ..", args{"$ cd .."}, cd, ".."},
		{"cd A", args{"$ cd A"}, cd, "A"},
		{"ls ", args{"$ ls"}, ls, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotState, gotArgument := getCommand(tt.args.line)
			if gotState != tt.wantState {
				t.Errorf("getCommand() gotState = %v, want %v", gotState, tt.wantState)
			}
			if gotArgument != tt.wantArgument {
				t.Errorf("getCommand() gotArgument = %v, want %v", gotArgument, tt.wantArgument)
			}
		})
	}
}

/*
dir a
14848514 b.txt
8504156 c.dat
*/
func Test_getFile(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
		wantSize int
	}{
		{"dir A", args{"dir A"}, "A", 0},
		{"14848514 b.txt", args{"14848514 b.txt"}, "b.txt", 14848514},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotSize := getFile(tt.args.line)
			if gotName != tt.wantName {
				t.Errorf("getFile() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotSize != tt.wantSize {
				t.Errorf("getFile() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}
