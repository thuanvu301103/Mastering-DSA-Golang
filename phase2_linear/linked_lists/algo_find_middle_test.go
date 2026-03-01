package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_FindMiddle(t *testing.T) {
	is := assert.New(t)

	t.Run("Odd_Length_List", func(t *testing.T) {
		// List: 1 -> 2 -> 3 -> 4 -> 5
		// Middle should be: 3
		l := &SinglyLinkedList{}
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)
		l.PushBack(4)
		l.PushBack(5)

		mid, ok := l.FindMiddle()
		is.True(ok)
		is.Equal(3, mid.Value)
	})

	t.Run("Even_Length_List", func(t *testing.T) {
		// List: 1 -> 2 -> 3 -> 4 -> 5 -> 6
		// Slow pointer will land on: 4
		l := &SinglyLinkedList{}
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)
		l.PushBack(4)
		l.PushBack(5)
		l.PushBack(6)

		mid, ok := l.FindMiddle()
		is.True(ok)
		is.Equal(4, mid.Value)
	})

	t.Run("Single_Node_List", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(10)

		mid, ok := l.FindMiddle()
		is.True(ok)
		is.Equal(10, mid.Value)
	})

	t.Run("Empty_List", func(t *testing.T) {
		l := &SinglyLinkedList{}
		mid, ok := l.FindMiddle()
		is.False(ok)
		is.Nil(mid)
	})
}
