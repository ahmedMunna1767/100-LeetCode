package problems

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// Brute Force
func ShortestPalindrome(s string) string {
	if isPalindrome(s) {
		return s
	}

	for i := 0; i < len(s); i++ {
		checkStr := s
		for j := i; j >= 0; j-- {
			checkStr = string(s[len(s)-1-j]) + checkStr
		}
		if isPalindrome(checkStr) {
			return checkStr
		}
		// fmt.Println(checkStr)
	}
	return s
}

func TestSSS() {
	fmt.Println(ShortestPalindrome("aacecaaa"))
	fmt.Println(ShortestPalindrome("abcd"))
	fmt.Println(ShortestPalindrome("aabba"))
	fmt.Println(ShortestPalindrome("abeeb"))
}
