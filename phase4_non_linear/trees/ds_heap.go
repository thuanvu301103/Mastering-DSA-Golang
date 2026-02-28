package tree

// MaxHeap represents a binary max heap data structure
type MaxHeap struct {
	data []int
}

// Insert adds an element and bubbles it up to maintain the max-heap property
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

// ExtractMax removes and returns the maximum element (root)
// This version matches the sequential order test case
func (h *MaxHeap) ExtractMax() int {
	val, ok := h.ExtractMaxSafe()
	if !ok {
		panic("ExtractMax from empty heap")
	}
	return val
}

// ExtractMaxSafe provides a way to extract without panicking,
// returning a boolean to indicate success
func (h *MaxHeap) ExtractMaxSafe() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	max := h.data[0]
	lastIdx := len(h.data) - 1

	// Move the last element to the root
	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]

	if !h.IsEmpty() {
		h.bubbleDown(0)
	}

	return max, true
}

// Peek returns the maximum element without removing it
// The test expects this to panic on an empty heap
func (h *MaxHeap) Peek() int {
	if h.IsEmpty() {
		panic("Peeking empty heap")
	}
	return h.data[0]
}

// IsEmpty checks if the heap has no elements
func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// GetInternalSlice returns the underlying array for inspection
func (h *MaxHeap) GetInternalSlice() []int {
	return h.data
}

// Internal helper to restore max property by moving a node up
func (h *MaxHeap) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[index] <= h.data[parent] {
			break
		}
		// Swap with parent
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

// Internal helper to restore max property by moving a node down
func (h *MaxHeap) bubbleDown(index int) {
	lastIdx := len(h.data) - 1
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left <= lastIdx && h.data[left] > h.data[largest] {
			largest = left
		}
		if right <= lastIdx && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		// Swap with largest child
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}
