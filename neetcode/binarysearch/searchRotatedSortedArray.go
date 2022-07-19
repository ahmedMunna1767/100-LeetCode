package binarysearch

import "fmt"

func rotatedSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + right
		middle = middle / 2
		fmt.Println(left, middle, right)
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			// left portion
			if nums[middle] < target || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			// right portion
			if target < nums[middle] || target > nums[right] {
				right = middle - 1

			} else {

				left = middle + 1
			}
		}
	}
	return -1
}

func RotatedSearch() {
	println(rotatedSearch([]int{5, 1, 3}, 5))
}
