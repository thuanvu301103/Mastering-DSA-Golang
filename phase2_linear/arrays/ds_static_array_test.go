package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticArrayOps(t *testing.T) {
	is := assert.New(t)
	// Initialize the implementation
	ops := StaticArrayOps{}

	t.Run("SumArray", func(t *testing.T) {
		input := [5]int{10, 20, 30, 40, 50}
		is.Equal(150, ops.SumArray(input))
	})

	t.Run("AddMatrices", func(t *testing.T) {
		m1 := [2][2]int{{1, 1}, {1, 1}}
		m2 := [2][2]int{{2, 2}, {2, 2}}
		expected := [2][2]int{{3, 3}, {3, 3}}
		is.Equal(expected, ops.AddMatrices(m1, m2))
	})

	t.Run("CharFrequency", func(t *testing.T) {
		is.Equal(3, ops.GetCharFrequency("banana")['a'-'a'])
	})

	t.Run("PointerUpdate", func(t *testing.T) {
		myArr := [3]int{1, 2, 3}
		ops.UpdateFirstElement(&myArr, 999)
		is.Equal(999, myArr[0])
	})
}
