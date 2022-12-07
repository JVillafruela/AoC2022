package day05

import (
	"reflect"
	"testing"
)

func TestNewRuneStack(t *testing.T) {
	got := NewRuneStack(3)
	want := RuneStack{4, 0, []rune{0, 0, 0}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewRuneStack got %v want %v", got, want)
	}
}

func TestRuneStack_isEmptyXXX(t *testing.T) {
	s := NewRuneStack(3)
	got := s.isEmpty()
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRuneStack_isEmpty(t *testing.T) {
	empty := NewRuneStack(3)
	stack := NewRuneStack(3)
	stack.pushOne('A')

	tests := []struct {
		name string
		s    *RuneStack
		want bool
	}{
		{"empty", &empty, true},
		{"not empty", &stack, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isEmpty(); got != tt.want {
				t.Errorf("RuneStack.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneStack_isFull(t *testing.T) {
	empty := NewRuneStack(2)
	stack1 := RuneStack{2, 1, []rune{'A', 0}}
	stack2 := RuneStack{2, 2, []rune{'A', 'B'}}

	tests := []struct {
		name string
		s    *RuneStack
		want bool
	}{
		{"empty", &empty, false},
		{"one", &stack1, false},
		{"full", &stack2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isFull(); got != tt.want {
				t.Errorf("RuneStack.isFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneStack_pushOne(t *testing.T) {

	s := NewRuneStack(2)
	s.pushOne('A')
	want := RuneStack{2, 1, []rune{'A', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(A) : got %v want %v", s, want)
	}

	s.pushOne('B')
	want = RuneStack{2, 2, []rune{'A', 'B'}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(B) : got %v want %v", s, want)
	}

	s.pushOne('C')
	want = RuneStack{3, 3, []rune{'A', 'B', 'C'}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(B) : got %v want %v", s, want)
	}
}

func TestRuneStack_popOne(t *testing.T) {
	s := RuneStack{2, 2, []rune{'A', 'B'}}
	r, err := s.popOne()
	if err != nil {
		t.Errorf("popOne(A,B) : error, stack not empty")
	}
	if r != 'B' {
		t.Errorf("popOne(A,B) : error got %c want %c ", r, 'B')
	}
	want := RuneStack{2, 1, []rune{'A', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(B) : got %v want %v", s, want)
	}

	r, err = s.popOne()
	if err != nil {
		t.Errorf("popOne(A) : error, stack not empty")
	}
	if r != 'A' {
		t.Errorf("popOne(A) : error got %c want %c ", r, 'A')
	}
	want = RuneStack{2, 0, []rune{0, 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(B) : got %v want %v", s, want)
	}

	r, err = s.popOne()
	if err == nil {
		t.Errorf("popOne(empty) : error, stack empty")
	}
	if r != 0 {
		t.Errorf("popOne(empty) : error got %c want %c ", r, 0)
	}
	want = RuneStack{2, 0, []rune{0, 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(B) : got %v want %v", s, want)
	}
}

func TestRuneStack_pushMany(t *testing.T) {
	s := NewRuneStack(3)
	s.pushMany([]rune{'A', 'B'})
	want := RuneStack{3, 2, []rune{'A', 'B', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(A) : got %v want %v", s, want)
	}

	s.pushMany([]rune{'C', 'D'})
	want = RuneStack{4, 4, []rune{'A', 'B', 'C', 'D'}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("pushone(A) : got %v want %v", s, want)
	}

}

func TestRuneStack_popMany(t *testing.T) {

	stack := RuneStack{3, 3, []rune{'A', 'B', 'C'}}

	type args struct {
		count int
	}
	tests := []struct {
		name    string
		s       *RuneStack
		args    args
		want    []rune
		wantErr bool
	}{
		{"pop 1", &stack, args{1}, []rune{'C'}, false},
		{"pop 2", &stack, args{2}, []rune{'B', 'A'}, false},
		{"pop 3", &stack, args{3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.popMany(tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("RuneStack.popMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RuneStack.popMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneStack_poke(t *testing.T) {
	s := NewRuneStack(3)
	err := s.poke(0, 'X')
	if err != nil {
		t.Errorf("poke(X,0) : error, stack not full")
	}
	want := RuneStack{3, 1, []rune{'X', 0, 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("poke(B,0) : got %v want %v", s, want)
	}

	err = s.poke(1, 'A')
	if err != nil {
		t.Errorf("poke(A,1) : error, stack not full")
	}
	want = RuneStack{3, 2, []rune{'X', 'A', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("poke(A,1) : got %v want %v", s, want)
	}

	err = s.poke(0, 'B')
	if err != nil {
		t.Errorf("poke(B,0) : error, stack not full")
	}
	want = RuneStack{3, 2, []rune{'B', 'A', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("poke(B,0) : got %v want %v", s, want)
	}

	err = s.poke(4, 'C')
	if err == nil {
		t.Errorf("poke(C,4) : error, beyond capacity")
	}
	want = RuneStack{3, 2, []rune{'B', 'A', 0}}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("poke(C,4) : got %v want %v", s, want)
	}
}

func TestRuneStack_top(t *testing.T) {
	empty := NewRuneStack(2)
	stack1 := RuneStack{2, 1, []rune{'A', 0}}
	stack2 := RuneStack{2, 2, []rune{'A', 'B'}}

	tests := []struct {
		name string
		s    *RuneStack
		want rune
	}{
		{"empty", &empty, 0},
		{"one ", &stack1, 'A'},
		{"full", &stack2, 'B'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.top(); got != tt.want {
				t.Errorf("RuneStack.top() = %v, want %v", got, tt.want)
			}
		})
	}
}
