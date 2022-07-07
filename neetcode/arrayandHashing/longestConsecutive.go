package arrayandhashing

func LongestConsecutive(nums []int) int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			hash[nums[i]] = hash[nums[i]] + 1
		} else {
			hash[nums[i]] = 1
		}
	}

	longest := 0

	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]-1]; ok {
			continue
		}
		length := 1
		_, ok := hash[nums[i]+1]
		for ok {
			length += 1
			_, ok = hash[nums[i]+1]
		}
		if length > longest {
			longest = length
		}
	}
	return 0

}
