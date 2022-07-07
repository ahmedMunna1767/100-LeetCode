package twopointers

func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		curSum := numbers[l] + numbers[r]

		if curSum > target {
			r -= 1

		} else if curSum < target {

			l += 1

		} else {
			return []int{l + 1, r + 1}

		}
	}

	return []int{}
}
