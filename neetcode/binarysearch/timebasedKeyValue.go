package binarysearch

import "fmt"

type valStruct struct {
	value     string
	timestamp int
}

type TimeMap struct {
	Value map[string][]valStruct
	/* PrevVal map[string]int */
}

func Constructor() TimeMap {
	val := make(map[string][]valStruct)
	return TimeMap{Value: val}
	/* prevVal := make(map[string]int)
	return TimeMap{Value: val, PrevVal: prevVal} */
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	/* if _, ok := tm.Value[key]; ok {
		tm.PrevVal[key] = timestamp
	} */
	tm.Value[key] = append(tm.Value[key], valStruct{value: value, timestamp: timestamp})
	fmt.Printf("%v\n", tm.Value)
}

/* func (tm *TimeMap) Get(key string, timestamp int) string {
	if timestamp < tm.PrevVal[key] {
		return ""
	}

	var largestPrev string
	largestPrev = ""
	for _, v := range tm.Value[key] {
		if v.timestamp == timestamp {
			largestPrev = v.value
			break
		} else if v.timestamp < timestamp {
			break
		}
		largestPrev = v.value
	}
	return largestPrev
}
*/

// Binary Search Solution
func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.Value[key]; !ok {
		return ""
	}
	values := tm.Value[key]
	left, right := 0, len(values)-1
	res := ""
	for left <= right {
		mid := (left + right) / 2
		if values[mid].timestamp <= timestamp {
			res = values[mid].value
			left = mid + 1

		} else {
			right = mid - 1
		}
	}
	return res
}
