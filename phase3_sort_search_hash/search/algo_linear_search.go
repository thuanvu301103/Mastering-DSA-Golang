package search

// LinearSearch looks for a target value in a slice.
// It returns the index of the first occurrence, or -1 if not found.
func LinearSearch(data []int, target int) int {
	for i, value := range data {
		// Check if current element matches target
		if value == target {
			return i // Found: return the current index
		}
	}
	return -1 // Not found: exhausted all elements
}