package slidingwindow

func LengthOfLongestSubstring(s string) int {
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
