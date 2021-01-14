package slidingwindow

func MaxSum(size int, arr []int) int {
	var result int
	for i := 0; i <= len(arr)-size; i++ {
		var sum int
		for j := i; j < i+size; j++ {
			sum += arr[j]
		}
		if i == 0 {
			result = sum
		}
		if result < sum {
			result = sum
		}

	}

	return result
}

func MaxSumSliding(size int, arr []int) int {
	var windowSum, maxSum, windowStart int
	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]
		if i >= size-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return maxSum
}
