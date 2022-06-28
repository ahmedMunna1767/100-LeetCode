package main

import (
	"fmt"

	"github.com/ahmedMunna1767/100-LeetCode/myHeap"
	"github.com/ahmedMunna1767/100-LeetCode/problems"
)

const (
	NOTVISITED = iota
	TRUE
	FALSE
)

// Problem One
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1Index := 0
	nums2Index := 0
	newArr := make([]int, 0)
	for nums1Index < len(nums1) && nums2Index < len(nums2) {
		if nums1[nums1Index] <= nums2[nums2Index] {
			newArr = append(newArr, nums1[nums1Index])
			nums1Index = nums1Index + 1
		} else {
			newArr = append(newArr, nums2[nums2Index])
			nums2Index = nums2Index + 1
		}
	}

	if nums1Index < len(nums1) {
		newArr = append(newArr, nums1[nums1Index:]...)
	}

	if nums2Index < len(nums2) {
		newArr = append(newArr, nums2[nums2Index:]...)
	}

	if len(newArr)%2 == 1 {
		return float64(newArr[len(newArr)/2])
	} else {
		retVal := (float64(newArr[len(newArr)/2]) + float64(newArr[(len(newArr)/2)-1])) / 2
		return float64(retVal)
	}
}

// Problem Two
// * Recursive Solution * //
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	// check if first character matches
	firstMatch := (!(s == "") && (p[0] == s[0] || p[0] == '.'))

	if len(p) >= 2 && p[1] == '*' {
		// if there is a klene star
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		// If 2nd character is not a klene start
		return firstMatch && (isMatch(s[1:], p[1:]))
	}
}

// * DP solution
func isMatchDynamic(s string, p string) bool {

	// Init the memoization array
	memo := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		memo[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			memo[i][j] = NOTVISITED
		}
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// if the node is previously computed, return appropriate value
		if memo[i][j] != NOTVISITED {
			return TRUE == memo[i][j]
		}

		var ans bool
		if j == len(p) {
			// reach both end at the same time
			ans = (i == len(s))
		} else {
			// if the first characters match
			firstMatch := (i < len(s) && (p[j] == s[i] || p[j] == '.'))

			// in case of a kleene star
			if j+1 < len(p) && p[j+1] == '*' {
				ans = (dp(i, j+2) || firstMatch && dp(i+1, j))
			} else {
				ans = firstMatch && dp(i+1, j+1)
			}
		}
		if ans {
			memo[i][j] = TRUE
		} else {
			memo[i][j] = FALSE
		}
		return ans
	}

	return dp(0, 0)
}

func testMergeKSortedList() {
	listOne := problems.ListNode{
		Val: 1,
		Next: &problems.ListNode{
			Val: 4,
			Next: &problems.ListNode{
				Val: 5,
			},
		},
	}

	listTwo := problems.ListNode{
		Val: 1,
		Next: &problems.ListNode{
			Val: 3,
			Next: &problems.ListNode{
				Val: 4,
			},
		},
	}

	listThree := problems.ListNode{
		Val: 2,
		Next: &problems.ListNode{
			Val: 6,
		},
	}

	node := (problems.MergeKSortedLists([]*problems.ListNode{&listOne, &listTwo, &listThree}))

	for node != nil {
		fmt.Print(node.Val, ", ")
		node = node.Next
	}

	fmt.Println()
}

func testFindKMedian() {
	fmt.Println(isMatchDynamic("aa", ".."))
	fmt.Println(isMatch("aa", ".."))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func main() {

	testMergeKSortedList()
	testFindKMedian()

	h := &myHeap.IntHeap{3, 5, 1, 2, 8, 4}
	myHeap.Init(h)
	fmt.Printf("%v", *h)
}
