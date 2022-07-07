package arrayandhashing

import "fmt"

func ContainsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := hashmap[nums[i]]; ok {
			return false
		} else {
			hashmap[nums[i]] = 1
		}
	}

	return true
}

func TestContainsDuplicate() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4, 5}))
}
