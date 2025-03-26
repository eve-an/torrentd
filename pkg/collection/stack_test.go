package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack[int](3)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	num, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, num)

	num, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, num)

	num, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, num)

	num, err = stack.Pop()
	assert.ErrorIs(t, err, ErrEmptyStack)
	assert.Equal(t, 0, num)
}
