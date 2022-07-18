package stack

import "fmt"

/* 40, 91, 123 */

func isValid(s string) bool {
	charStack := make([]byte, 0)
	open := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			charStack = append(charStack, s[i])
			open++
		} else {
			if open == 0 {
				return false
			}
			last := charStack[len(charStack)-1]
			if last+1 == s[i] || last+2 == s[i] {
				charStack = charStack[0 : len(charStack)-1]
				open--
			} else {
				return false
			}
		}
	}
	return len(charStack) == 0
}

func TestValidParenthesis() {
	fmt.Println(isValid("{[]}"))
}
