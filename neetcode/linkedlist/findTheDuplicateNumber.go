package linkedlist

import "fmt"

func findDuplicate(nums []int) int {
	highest := len(nums)
	total, expectedTotal := 0, 0
	for i, v := range nums {
		total += v
		expectedTotal += i + 1
	}

	// fmt.Println(total, expectedTotal)
	return highest - expectedTotal + total
}

func FindDuplicate() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
