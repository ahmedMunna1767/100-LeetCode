package twopointers

import "sort"

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		threeSum := 0
		for l < r {
			threeSum = a + nums[l] + nums[r]
			if threeSum > 0 {
				r -= 1
			} else if threeSum < 0 {
				l += 1

			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}

		}

	}

	return res
}
