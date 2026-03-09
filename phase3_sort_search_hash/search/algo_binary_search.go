package search

// BinarySearch performs a search on a SORTED slice.
// Returns the index of the target, or -1 if not found.
func BinarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		// Calculate middle index (avoiding overflow)
		mid := low + (high-low)/2

		if data[mid] == target {
			return mid // Found the target
		}

		if data[mid] < target {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}
