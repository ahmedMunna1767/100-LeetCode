package twopointers

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	newstr := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			newstr += string(s[i])
		}
	}
	fmt.Println(newstr)

	start, end := 0, len(newstr)-1

	for start < end {
		if newstr[start] != newstr[end] {
			return false
		}
		start++
		end--
	}

	return true
}
