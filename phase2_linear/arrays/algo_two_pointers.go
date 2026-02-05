package arrays

import "strings"

type TwoPointerOps struct{}

func (o TwoPointerOps) TwoSumSorted(sortedArr []int, target int) []int {
	var left = 0
	var right = len(sortedArr) - 1

	for left < right {
		sum := sortedArr[left] + sortedArr[right]
		if sum == target {
			break
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}

	return []int{left, right}
}

func (o TwoPointerOps) IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	var isNotPalindrome = false
	for i := 0; i < (len(str) / 2); i++ {
		if str[i] != str[len(str)-1-i] {
			isNotPalindrome = true
			break
		}
	}

	return !isNotPalindrome
}

func (o TwoPointerOps) Reverse(arr []int) {
	left := 0
	right := len(arr) - 1

	stillSwapping := true

	for stillSwapping && left < right {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--

		if left >= right {
			stillSwapping = false
		}
	}
}
