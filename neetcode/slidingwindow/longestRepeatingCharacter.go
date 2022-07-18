package slidingwindow

func CharacterReplacement(s string, k int) int {
	hash := make([]int, 26)

	l := 0
	maxf := 0
	res := 0
	for r := 0; r < len(s); r++ {
		hash[s[r]-'A'] = 1 + hash[s[r]-'A']
		maxf = Max(maxf, hash[s[r]-'A'])
		if (r-l+1)-maxf > k {
			hash[s[l]-'A'] -= 1
			l += 1

		}
		res = Max(res, r-l+1)

	}

	return res
}

func TestCR() {
	println(CharacterReplacement("ABAB", 2))
}
