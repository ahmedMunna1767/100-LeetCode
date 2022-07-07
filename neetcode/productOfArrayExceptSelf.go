package neetcode

func ProductExceptSelf(nums []int) []int {
	postFix := make([]int, len(nums))
	preFix := make([]int, len(nums))
	retVal := make([]int, len(nums))

	preFix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preFix[i] = preFix[i-1] * nums[i]
	}

	postFix[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i++ {
		postFix[i] = postFix[i+1] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			retVal[0] = postFix[1]
		} else if i == len(nums)-1 {
			retVal[len(nums)-1] = preFix[len(nums)-1]
		} else {
			retVal[i] = preFix[i-1] * postFix[i+1]
		}
	}

	return retVal
}
