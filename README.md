# Mastering-DSA-Golang
A comprehensive, phase-based repository designed to master Data Structures and Algorithms using Golang. This project is structured as a professional learning roadmap, moving from mathematical foundations to specialist-level optimization.

## Preparation

1. Initialize the module (if you haven't yet):
```Bash
go mod init go-dsa-mastery
```

2. Install `Testify`
```Bash
go get github.com/stretchr/testify
```

3. Execute the tests:
```Bash
go test -v ./phase2_linear/arrays/
```

## Detailed Repository Layout
Each directory will contain the implementation file (the "how it works") and a corresponding test file (the "validation").

```Plaintext
root/
├── phase2_linear/
│   ├── arrays/
│   │   ├── ds_static_array.go
│   │   ├── ds_static_array_test.go
│   │   ├── ds_dynamic_array.go             <-- Structure/Logic
│   │   ├── ds_dynamic_array_test.go        <-- Test/Verification
│   │   ├── algo_two_pointers.go            <-- Algorithm
│   │   ├── algo_two_pointers_test.go       <-- Test/Verification
│   │ 	├── algo_sliding_window.go       
│   │   ├── algo_sliding_window_test.go
│   │ 	├── algo_prefix_sum.go       
│   │   └── algo_prefix_sum_test.go
│   ├── linked_lists/
│   │   ├── ds_singly_list.go               
│   │   ├── ds_singly_list_test.go          
│   │   ├── algo_cycle_detection.go         
│   │   └── algo_cycle_detection_test.go
│   ├── stacks_queues/
│   │   ├── ds_stack.go
│   │   ├── ds_stack_test.go               
│   │   ├── ds_queue.go          
│   │   └── ds_queue_test.go    
│   └── README.md                           <-- Theory & Documentation
├── phase3_sort_search_hash/
│   ├── sort/
│   │   ├── ds_max_heap.go
│   │   ├── ds_max_heap_test.go
│   │   ├── algo_bubble_sort.go
│   │   ├── algo_bubble_sort_test.go
│   │   ├── algo_insertion_sort.go
│   │   ├── algo_insertion_sort_test.go
│   │   ├── algo_selection_sort.go
│   │   ├── algo_selection_sort_test.go
│   │   ├── algo_merge_sort.go
│   │   ├── algo_merge_sort_test.go
│   │   ├── algo_quick_sort.go
│   │   ├── algo_quick_sort_test.go
│   │   ├── algo_heap_sort.go
│   │   └── algo_heap_test.go
```

## File Content Standards
To show the difference clearly, here is how you should structure the code within these files using Generics.

### The Structure File (.go)
This file defines the data type and its methods. It should be clean and well-commented.

File: `phase2_linear/arrays/ds_static_arsay.go`

```Go
package arrays

// StaticArrayOps is a concrete implementation of the ArrayOps interface
type StaticArrayOps struct{}

// SumArray implements ArrayOps
func (s StaticArrayOps) SumArray(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// AddMatrices implements ArrayOps
func (s StaticArrayOps) AddMatrices(a, b [2][2]int) [2][2]int {
	var result [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

// GetCharFrequency implements ArrayOps
func (s StaticArrayOps) GetCharFrequency(str string) [26]int {
	var freq [26]int
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			freq[char-'a']++
		}
	}
	return freq
}

// UpdateFirstElement implements ArrayOps
func (s StaticArrayOps) UpdateFirstElement(arr *[3]int, newValue int) {
	if arr != nil {
		arr[0] = newValue
	}
}

```

### The Test File (_test.go)
This file uses Go's built-in testing package to verify the logic.

File: `phase2_linear/arrays/ds_static_arrays_test.go`

```Go
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
```