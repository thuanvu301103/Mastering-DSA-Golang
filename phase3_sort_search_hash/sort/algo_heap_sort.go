package sort

func HeapSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	// 1. Build Max Heap (Rearrange array)
	// We start from the last non-leaf node: index = (n/2) - 1
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	// 2. Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root (largest) to the end
		array[0], array[i] = array[i], array[0]

		// Call max heapify on the reduced heap
		heapify(array, i, 0)
	}
}

// heapify maintains the max heap property for a subtree rooted at index i
func heapify(array []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// If left child is larger than root
	if left < n && array[left] > array[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && array[right] > array[largest] {
		largest = right
	}

	// If largest is not root, swap and continue heapifying
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}
