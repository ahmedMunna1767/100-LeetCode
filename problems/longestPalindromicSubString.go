package problems

import "fmt"

func ifPalindrome(left, right int, s string) bool {
	if left < 0 || right >= len(s) {
		return false
	}
	return s[left] == s[right]
}

func LongestPalindromicSubString(s string) string {
	maxStr := ""
	lenMaxStr := 0
	for i := 0; i < len(s); i++ {
		right := i + 1
		left := i - 1
		str := string(s[i])
		for {
			palindrome := ifPalindrome(left, right, s)
			if !palindrome {
				break
			}
			str = string(s[left]) + str + string(s[right])
			right++
			left--
		}
		if lenMaxStr < len(str) {
			lenMaxStr = len(str)
			maxStr = str
		}

		right = i + 1
		left = i
		str = ""
		for {
			palindrome := ifPalindrome(left, right, s)
			if !palindrome {
				break
			}
			str = string(s[left]) + str + string(s[right])
			right++
			left--
		}
		if lenMaxStr < len(str) {
			lenMaxStr = len(str)
			maxStr = str
		}
	}
	return maxStr
}

func TestLPSS() {
	fmt.Println(LongestPalindromicSubString("babab"))
	fmt.Println(LongestPalindromicSubString("cbbd"))
}
