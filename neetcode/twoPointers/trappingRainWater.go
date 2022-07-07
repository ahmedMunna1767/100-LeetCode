package twopointers

func TrapOne(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	minLR := make([]int, len(height))

	maxLeft[0] = 0
	maxRight[len(height)-1] = 0

	result := 0

	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	for i := len(height) - 1; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	for i := 0; i < len(height); i++ {
		minLR[i] = min(maxLeft[i], maxRight[i])
	}

	for i := 0; i < len(height); i++ {
		result = result + max(0, minLR[i]-height[i])
	}

	return result
}

func TrapTwo(height []int) int {
	result := 0

	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			result = result + maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			result = result + maxRight - height[right]
		}
	}

	return result
}
