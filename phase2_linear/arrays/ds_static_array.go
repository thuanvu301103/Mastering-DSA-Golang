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
