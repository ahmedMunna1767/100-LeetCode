package arrayandhashing

import (
	"fmt"
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	reverseHash := make(map[int][]int, 0)
	retVal := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			hash[nums[i]] = hash[nums[i]] + 1
		} else {
			hash[nums[i]] = 1
		}
	}
	fmt.Printf("%v\n", hash)

	keys := make([]int, 0)
	for key, v := range hash {
		reverseHash[v] = append(reverseHash[v], key)
	}

	for key := range reverseHash {
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Printf("%v  %v\n", reverseHash, keys)

	i := 0
	for len(retVal) < k {
		retVal = append(retVal, reverseHash[keys[i]]...)
		i = i + 1
	}

	return retVal
}

func RunTopKFreqElem() {
	fmt.Println(TopKFrequent([]int{5, 5, 5, 2, 2, 2, 4, 3}, 2))
}
