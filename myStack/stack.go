package myStack

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Length() int {
	return len(s.vals)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length() == 0 {
		var zero T
		return zero, false
	}
	top := s.Top()
	s.vals = s.vals[:s.Length()-1]
	return top, true
}

func (s *Stack[T]) Bottom() T {
	return s.vals[0]
}

func (s *Stack[T]) Top() T {
	return s.vals[len(s.vals)-1]
}

func CheckStack() {
	stack := Stack[int]{}

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for stack.Length() > 0 {
		fmt.Println(stack.Pop())
	}
}
