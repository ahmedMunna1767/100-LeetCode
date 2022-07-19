package binarysearch

import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	max := math.MinInt
	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	left, right := 1, max
	retVal := 0
	for left <= right {
		middle := (left + right) / 2
		totaltime := 0
		for _, v := range piles {
			totaltime += ((v - 1) / middle) + 1
		}

		if totaltime <= h {
			retVal = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return retVal
}

func MinEatingSpeed() {
	println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
