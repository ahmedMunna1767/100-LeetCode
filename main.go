package main

import "fmt"

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

func main() {
	fmt.Println("Hi.....")

	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
