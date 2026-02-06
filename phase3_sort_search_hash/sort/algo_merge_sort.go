package sort

func MergeSort(array []int) []int {
	// Base case: 0 or 1 elements are already "sorted"
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	leftSortedArray := MergeSort(array[:mid])
	rightSortedArray := MergeSort(array[mid:])

	sortedArray := make([]int, 0, len(array))

	for len(leftSortedArray) > 0 && len(rightSortedArray) > 0 {
		if leftSortedArray[0] <= rightSortedArray[0] {
			sortedArray = append(sortedArray, leftSortedArray[0])
			leftSortedArray = leftSortedArray[1:]
		} else {
			sortedArray = append(sortedArray, rightSortedArray[0])
			rightSortedArray = rightSortedArray[1:]
		}
	}

	sortedArray = append(sortedArray, leftSortedArray...)
	sortedArray = append(sortedArray, rightSortedArray...)

	return sortedArray
}
