package day05

import "errors"

/*
type Stack interface {
	pushOne(item any) error
	pushMany(items []any) error
	popOne() (any, error)
	popMany(count int) ([]any, error)
	//peek() any
	poke(position int, item any)
	top() any
	isFull() bool
	isEmpty() bool
}
*/
type RuneStack struct {
	capacity int
	count    int
	items    []rune
}

func NewRuneStack(max int) RuneStack {
	var s RuneStack
	s.capacity = max
	s.items = make([]rune, s.capacity)
	return s
}

func (s *RuneStack) isEmpty() bool {
	return s.count == 0
}

func (s *RuneStack) isFull() bool {
	return s.count == s.capacity
}

func (s *RuneStack) pushOne(r rune) {
	if s.isFull() {
		s.items = append(s.items, r)
		s.count++
		s.capacity++
		return
	}
	s.items[s.count] = r
	s.count++
}

func (s *RuneStack) popOne() (rune, error) {
	if s.isEmpty() {
		return rune(0), errors.New("stack empty")
	}
	s.count--
	r := s.items[s.count]
	s.items[s.count] = rune(0)
	return r, nil
}

func (s *RuneStack) popMany(count int) ([]rune, error) {
	if s.isEmpty() {
		return nil, errors.New("stack empty")
	}

	var items []rune
	for i := 0; i < count; i++ {
		r, err := s.popOne()
		if err != nil {
			return items, err
		}
		items = append(items, r)
	}
	return items, nil
}

func (s *RuneStack) pushMany(items []rune) {
	for _, r := range items {
		s.pushOne(r)
	}
}

// memory of the TRS-80 BASIC ;-)
func (s *RuneStack) poke(position int, r rune) error {
	if position+1 > s.capacity {
		return errors.New("poke beyond stack capacity")
	}
	if position >= s.count {
		s.count = position + 1
	}
	s.items[position] = r
	return nil
}

// peek at the top of stack
func (s *RuneStack) top() rune {
	if s.count == 0 {
		return 0
	}
	return s.items[s.count-1]
}
