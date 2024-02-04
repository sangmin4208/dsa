package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntStack(t *testing.T) {
	stack := New[int]()
	t.Run("should create an empty stack", func(t *testing.T) {
		assert.Equal(t, 0, stack.Size())
		assert.True(t, stack.IsEmpty())
	})

	t.Run("should push elements to the stack", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Size())
		assert.False(t, stack.IsEmpty())
	})

	t.Run("should pop elements from the stack", func(t *testing.T) {
		assert.Equal(t, 2, stack.Pop())
		assert.Equal(t, 1, stack.Pop())
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("should peek the top element of the stack", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Peek())
		stack.Pop()
		assert.Equal(t, 1, stack.Peek())
	})
}
