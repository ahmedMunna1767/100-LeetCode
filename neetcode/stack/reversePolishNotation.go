package stack

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	charStack := make([]int, 0)
	for _, v := range tokens {
		fmt.Printf("%v", charStack)
		println(v)
		if v == "+" || v == "-" || v == "*" || v == "/" {
			b := charStack[len(charStack)-1]
			charStack = charStack[:len(charStack)-1]
			if v == "+" {
				charStack[len(charStack)-1] = charStack[len(charStack)-1] + b
			} else if v == "-" {
				charStack[len(charStack)-1] = charStack[len(charStack)-1] - b

			} else if v == "*" {
				charStack[len(charStack)-1] = charStack[len(charStack)-1] * b

			} else {
				charStack[len(charStack)-1] = charStack[len(charStack)-1] / b
			}
		} else {
			intVal, err := strconv.Atoi(v)
			if err != nil {
				return 0
			}
			charStack = append(charStack, intVal)
		}
	}
	fmt.Printf("%v\n", charStack)
	return charStack[0]
}

func EvalRPN() {
	println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
