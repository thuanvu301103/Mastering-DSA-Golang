package arrays

type SlidingWindowOps struct{}

func (o SlidingWindowOps) MaxSumFixed(nums []int, k int) int {
	n := len(nums)
	if n < k || k <= 0 {
		return 0
	}

	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k; i < n; i++ {
		currentSum = currentSum + nums[i] - nums[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func (o SlidingWindowOps) FindAverages(nums []int, k int) []float64 {
	n := len(nums)
	if n < k || k <= 0 {
		return []float64{}
	}

	result := make([]float64, n-k+1)
	windowSum := 0.0
	windowStart := 0

	for windowEnd := 0; windowEnd < n; windowEnd++ {
		windowSum += float64(nums[windowEnd])

		if windowEnd >= k-1 {
			result[windowStart] = windowSum / float64(k)

			windowSum -= float64(nums[windowStart])
			windowStart++
		}
	}

	return result
}

func (o SlidingWindowOps) MinSizeSubarraySum(target int, nums []int) int {
	windowSum := 0
	windowStart := 0
	minLength := len(nums) + 1

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		// 1. STRETCH
		windowSum += nums[windowEnd]

		// 2. SHRINK
		for windowSum >= target {
			currentWindowLen := windowEnd - windowStart + 1

			if currentWindowLen < minLength {
				minLength = currentWindowLen
			}

			windowSum -= nums[windowStart]
			windowStart++
		}
	}

	if minLength > len(nums) {
		return 0
	}

	return minLength
}

func (o SlidingWindowOps) LongestSubstringKDistinct(str string, k int) int {
	if k == 0 || len(str) == 0 {
		return 0
	}

	windowStart := 0
	maxLength := 0
	charFrequency := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		charFrequency[rightChar]++

		// SHRINK
		for len(charFrequency) > k {
			leftChar := str[windowStart]
			charFrequency[leftChar]--

			if charFrequency[leftChar] == 0 {
				delete(charFrequency, leftChar)
			}
			windowStart++
		}

		currentWindowLen := windowEnd - windowStart + 1
		if currentWindowLen > maxLength {
			maxLength = currentWindowLen
		}
	}

	return maxLength
}
