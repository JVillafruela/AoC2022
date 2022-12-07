package day05

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	s := []rune{}
	reverse(s)
	want := []rune{}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("reveser() : got %v want %v", s, want)
	}

	s = []rune{'A'}
	reverse(s)
	want = []rune{'A'}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("reveser(A) : got %v want %v", s, want)
	}

	s = []rune{'A', 'B'}
	reverse(s)
	want = []rune{'B', 'A'}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("reveser(A,B) : got %v want %v", s, want)
	}

	s = []rune{'A', 'B', 'C'}
	reverse(s)
	want = []rune{'C', 'B', 'A'}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("reveser(A,B,C) : got %v want %v", s, want)
	}

	s = []rune{'A', 'B', 'C', 'D'}
	reverse(s)
	want = []rune{'D', 'C', 'B', 'A'}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("reveser(A,B,C,D) : got %v want %v", s, want)
	}

}
