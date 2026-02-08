package sort

// QuickSort is the public wrapper function for the sorting operation.
// Users only need to pass the array, making it cleaner to use.
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	// Start the recursive process with the full range of the array
	quickSortRecursive(array, 0, len(array)-1)
	return array
}

// quickSortRecursive is the internal function that handles the divide and conquer logic.
func quickSortRecursive(array []int, low, high int) {
	if low < high {
		// p is the partitioning index, array[p] is now at the correct spot
		p := partition(array, low, high)

		// Recursively sort elements before and after partition
		quickSortRecursive(array, low, p-1)
		quickSortRecursive(array, p+1, high)
	}
}

// partition handles the core logic of picking a pivot and rearranging elements.
func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}
