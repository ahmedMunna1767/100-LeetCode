package neetcode

func TwoSum(nums []int, target int) []int {

	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]
		if _, ok := hashmap[need]; ok {
			return []int{hashmap[need], i}
		} else {
			hashmap[nums[i]] = i
		}
	}

	return []int{}
}
