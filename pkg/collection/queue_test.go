package collection_test

import (
	"testing"

	"github.com/eve-an/torrentd/pkg/collection"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := collection.NewQueue[int](3)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	num, err := queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, num)

	num, err = queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, num)

	num, err = queue.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, num)

	num, err = queue.Dequeue()
	assert.ErrorIs(t, err, collection.ErrEmptyQueue)
	assert.Equal(t, 0, num)
}
