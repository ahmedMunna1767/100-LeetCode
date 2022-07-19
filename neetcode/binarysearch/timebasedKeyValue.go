package binarysearch

import "fmt"

type valStruct struct {
	value     string
	timestamp int
}

type TimeMap struct {
	Value   map[string][]valStruct
	PrevVal map[string]int
}

func Constructor() TimeMap {
	val := make(map[string][]valStruct)
	prevVal := make(map[string]int)
	return TimeMap{Value: val, PrevVal: prevVal}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.Value[key]; ok {
		this.PrevVal[key] = timestamp
	}
	this.Value[key] = append(this.Value[key], valStruct{value: value, timestamp: timestamp})
	fmt.Printf("%v\n", this.Value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if timestamp < this.PrevVal[key] {
		return ""
	}

	var largestPrev string
	largestPrev = ""
	for _, v := range this.Value[key] {
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
