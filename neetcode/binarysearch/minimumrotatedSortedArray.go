package binarysearch

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := nums[0]

	for left <= right {
		if nums[left] < nums[right] {
			if nums[left] < res {
				res = nums[left]
			}
			break
		}
		middle := left + right
		middle = middle / 2
		fmt.Println(left, middle, right)
		if nums[middle] < res {
			res = nums[middle]
		}
		if nums[middle] >= nums[left] {
			// left portion
			left = middle + 1

		} else {
			// right portion
			right = middle - 1
		}
	}
	return res
}

func FindMin() {
	println(findMin([]int{3, 4, 5, 1, 2}))
}
