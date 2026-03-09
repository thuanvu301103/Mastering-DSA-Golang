package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTree() *Node {
	return &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left:  &Node{Value: 4},
			Right: &Node{Value: 5},
		},
		Right: &Node{Value: 3},
	}
}

func TestTreeTraversal(t *testing.T) {
	is := assert.New(t)
	root := setupTree()

	t.Run("InOrder_Traversal", func(t *testing.T) {
		// Expected: [4, 2, 5, 1, 3]
		var result []int
		InOrder(root, &result)
		is.Equal([]int{4, 2, 5, 1, 3}, result)
	})

	t.Run("PreOrder_Traversal", func(t *testing.T) {
		// Expected: [1, 2, 4, 5, 3]
		var result []int
		PreOrder(root, &result)
		is.Equal([]int{1, 2, 4, 5, 3}, result)
	})

	t.Run("PostOrder_Traversal", func(t *testing.T) {
		// Expected: [4, 5, 2, 3, 1]
		var result []int
		PostOrder(root, &result)
		is.Equal([]int{4, 5, 2, 3, 1}, result)
	})

	t.Run("Empty_Tree", func(t *testing.T) {
		var result []int
		InOrder(nil, &result)
		is.Empty(result, "Traversing nil should return empty slice")
	})
}
