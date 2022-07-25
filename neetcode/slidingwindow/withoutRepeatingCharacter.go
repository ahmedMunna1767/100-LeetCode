package slidingwindow

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]int, 0)
	l := 0
	res := 0

	for i := 0; i < len(s); i++ {
		for {
			_, ok := charSet[s[i]]
			if !ok {
				break
			}
			delete(charSet, s[l])
			l += 1
		}
		charSet[s[i]] = i
		res = Max(res, i-l+1)
	}
	return res
}

/* func lengthOfLongestSubstring(s string) int {

	charSet := make(map[byte]int, 0)
	slow := 0
	res := 0
	for fast := 0; fast < len(s); fast++ {
		prev, ok := charSet[s[fast]]
		if ok {
			if slow < prev+1 {
				slow = prev + 1
			} else {
				slow = fast
			}
		}
		thisLen := fast - slow + 1
		if thisLen > res {
			res = thisLen
		}
		charSet[s[fast]] = fast
	}

	return res
} */

func LengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcdebb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
