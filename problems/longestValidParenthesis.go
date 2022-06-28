package problems

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func LongestValidParenthesis(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	maxans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {

			if len(stack) == 1 {
				stack[0] = i
			} else {
				stack = stack[0 : len(stack)-1]
				maxans = max(maxans, i-stack[len(stack)-1])
			}
		}
		fmt.Printf("%v\n", stack)
	}
	return maxans
}

func TestLVP() {
	println(LongestValidParenthesis("(()"))
}
