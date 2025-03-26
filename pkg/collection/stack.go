package collection

import "errors"

var ErrEmptyStack = errors.New("stack is empty and cannot be popped")

type Stack[T any] struct {
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, size),
	}
}

func (s *Stack[T]) Push(t T) {
	s.data = append(s.data, t)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return *new(T), ErrEmptyStack
	}

	idx := len(s.data) - 1

	var t T
	t, s.data = s.data[idx], s.data[:idx]

	return t, nil
}
