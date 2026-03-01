package linked_list

import (
	"reflect"
	"testing"
)

// Helper function to convert a linked list to a slice for easy comparison
func listToSlice(l *SinglyLinkedList) []int {
	// Initialize with an empty slice instead of letting it be nil
	res := []int{}
	curr := l.Head
	for curr != nil {
		res = append(res, curr.Value)
		curr = curr.Next
	}
	return res
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal list",
			input:    []int{10, 20, 30, 40},
			expected: []int{40, 30, 20, 10},
		},
		{
			name:     "Single node list",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Two nodes",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the list
			list := &SinglyLinkedList{}
			// Push items (using a loop to build the list from the input slice)
			// Note: because Push/Prepend reverses order, we iterate backwards
			// to match the visual 'input' slice order
			for i := len(tt.input) - 1; i >= 0; i-- {
				list.Push(tt.input[i])
			}

			// Execute Reverse
			list.Reverse()

			// Get result as slice
			result := listToSlice(list)

			// Compare result with expected
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test '%s' failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
