package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	is := assert.New(t)

	t.Run("Singly_PushFront_Updates_Head", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushFront(10)
		l.PushFront(20)

		// In a singly list, new nodes become the new Head
		is.Equal(20, l.Head.Value)
		is.Equal(10, l.Head.Next.Value)
		is.Equal(2, l.Size)
	})

	t.Run("Singly_PushBack_Traverses_To_End", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(1)
		l.PushBack(2)

		is.Equal(1, l.Head.Value)
		is.Equal(2, l.Head.Next.Value)
		is.Nil(l.Head.Next.Next, "Last node's Next must be nil")
	})

	t.Run("Singly_Remove_Middle_Node", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)

		// Removing 20 should link 10 directly to 30
		success := l.Remove(20)
		is.True(success)
		is.Equal(2, l.Size)
		is.Equal(30, l.Head.Next.Value)
	})

	t.Run("Singly_Remove_Head_Node", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(5)
		l.PushBack(10)

		l.Remove(5)
		is.Equal(10, l.Head.Value, "Head should now be the second node")
		is.Equal(1, l.Size)
	})

	t.Run("Singly_Search_Value", func(t *testing.T) {
		l := &SinglyLinkedList{}
		l.PushBack(100)

		is.True(l.Find(100))
		is.False(l.Find(999))
	})

	t.Run("Singly_Empty_List_State", func(t *testing.T) {
		l := &SinglyLinkedList{}
		is.True(l.IsEmpty())
		is.False(l.Remove(10), "Should not be able to remove from empty list")
		is.Nil(l.Head)
	})
}
