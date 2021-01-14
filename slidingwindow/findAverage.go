package slidingwindow

func FindAverages(size int, arr []int) []float64 {
	result := make([]float64, len(arr)-size+1)
	for i := 0; i <= len(arr)-size; i++ {
		var sum float64
		for j := i; j < i+size; j++ {
			sum += float64(arr[j])
		}
		result[i] = sum / float64(size)
	}

	return result
}

func FindAveragesSliding(size int, arr []int) []float64 {
	result := make([]float64, len(arr)-size+1)
	var windowSum float64
	var windowStart int
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += float64(arr[windowEnd])
		if windowEnd >= size-1 {
			result[windowStart] = windowSum / float64(size)
			windowSum -= float64(arr[windowStart])
			windowStart++
		}
	}
	return result
}
