package twopointers

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxArea(height []int) int {
	maxArea := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(area, maxArea)
		if height[left] < height[right] {
			left += 1
		} else if height[right] <= height[left] {
			right -= 1
		}
	}

	return maxArea
}
