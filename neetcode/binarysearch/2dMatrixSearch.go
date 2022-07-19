package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	targetArr := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i][0] <= target {
			targetArr = i
			break
		}
	}

	return func(nums []int, target int) bool {
		left := 0
		right := len(nums) - 1

		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return true
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return false
	}(matrix[targetArr], target)
}

func SearchMatrix() {
	println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}
