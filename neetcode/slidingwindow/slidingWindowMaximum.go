package slidingwindow

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)

	findMax := func(nums []int) int {
		max := math.MinInt
		for _, v := range nums {
			if v > max {
				max = v
			}
		}
		return max
	}

	max := findMax(nums[0:k])
	result = append(result, max)

	fmt.Println(nums[0:k], max)

	for i := k; i < len(nums); i++ {
		if nums[i-k] == max {
			max = findMax(nums[i-k+1 : i+1])
		} else if nums[i] > max {
			max = nums[i]
		}

		fmt.Println(nums[i-k+1:i+1], max, nums[i-k])
		result = append(result, max)

	}

	return result
}

func MaxSlidingWindow() {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}
