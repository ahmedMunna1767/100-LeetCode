package kmp

import (
	"fmt"
)

func KmpStringSearch(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	kmpTable := BuildKmpTable(needle)
	i, j := 0, 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = kmpTable[j-1]
			}
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

func BuildKmpTable(p string) []int {
	kmpTable := make([]int, len(p))

	j := 0
	for i := 1; i < len(p); i++ {
		j = kmpTable[i-1]
		for j > 0 && p[i] != p[j] {
			j = kmpTable[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		kmpTable[i] = j
	}

	return kmpTable
}

func TestKMP() {
	fmt.Println(KmpStringSearch("abxabcabcaby", "abcaby"))
}
