package stack

import "math"

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	intstack := make([]int, 0)
	min := math.MaxInt

	return MinStack{
		stack: intstack,
		min:   min,
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if ms.min > val {
		ms.min = val
	}
}

func (ms *MinStack) Pop() {
	last := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if ms.min == last {
		ms.min = math.MaxInt
		for _, v := range ms.stack {
			if v < ms.min {
				ms.min = v
			}
		}
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min
}
