package slidingwindow

import "math"

func FindMinSubArray(number int, arr []int) int {
	var sum, windowStart int
	minLength := math.Inf(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for sum >= number {
			minLength = math.Min(minLength, float64(i-windowStart+1))
			sum -= arr[windowStart]
			windowStart++
		}
	}
	if math.IsInf(minLength, 0) {
		return 0
	}
	return int(minLength)
}
