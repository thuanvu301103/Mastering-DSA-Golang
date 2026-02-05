package sort

func InsertionSort(array []int) []int {

	for slit := 1; slit < len(array); slit++ {
		for i := slit - 1; i >= 0; i-- {
			if array[i+1] <= array[i] {
				array[i+1], array[i] = array[i], array[i+1]
			} else {
				break
			}
		}
	}

	return array
}
