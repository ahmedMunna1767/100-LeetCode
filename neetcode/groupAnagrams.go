package neetcode

import (
	"github.com/ahmedMunna1767/100-LeetCode/utility"
)

func GroupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	groups := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		sortedStr := utility.SortString(strs[i])
		hash[sortedStr] = append(hash[sortedStr], strs[i])
	}

	for _, value := range hash {
		groups = append(groups, value)
	}
	return groups
}
