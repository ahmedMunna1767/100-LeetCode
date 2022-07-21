package stack

import "fmt"

type tempIndex struct {
	temperature int
	index       int
}

/* func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	monotonicStack := make([]tempIndex, 0)

	for i := 0; i < len(temperatures)-1; i++ {

		j := i + 1
		for j < len(temperatures) {
			if temperatures[i] < temperatures[j] {
				break
			}
			j++
		}
		if j == len(temperatures) {
			result[i] = 0
		} else {
			result[i] = j - i
		}
	}

	return result
} */

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	monotonicStack := make([]tempIndex, 0)

	for i, v := range temperatures {
		for len(monotonicStack) != 0 && monotonicStack[len(monotonicStack)-1].temperature < v {
			ti := monotonicStack[len(monotonicStack)-1]
			monotonicStack = monotonicStack[:len(monotonicStack)-1]
			result[ti.index] = i - ti.index
		}
		monotonicStack = append(monotonicStack, tempIndex{temperature: v, index: i})
	}

	return result
}

func DailyTemperatures() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
