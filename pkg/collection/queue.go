package collection

import (
	"errors"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0, size),
	}
}

func (q *Queue[T]) Enqueue(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.data) == 0 {
		return *new(T), ErrEmptyQueue
	}

	var elem T
	elem, q.data = q.data[0], q.data[1:]

	return elem, nil
}
