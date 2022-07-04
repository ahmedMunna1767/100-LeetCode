package problems

import (
	"fmt"
	"sort"
)

func minDeletions(s string) int {
	freqArray := make([]int, 26)

	for i := 0; i < len(s); i++ {
		freqArray[s[i]-'a'] += 1
	}
	sort.Slice(freqArray, func(a, b int) bool {
		return freqArray[b] > freqArray[a]
	})

	deleteCount := 0
	// Maximum frequency the current character can have
	maxFreqAllowed := len(s)

	for i := 25; i >= 0; i-- {
		if freqArray[i] == 0 {
			break
		}
		if freqArray[i] > maxFreqAllowed {
			fmt.Println(freqArray[i], maxFreqAllowed)
			deleteCount += freqArray[i] - maxFreqAllowed
			freqArray[i] = maxFreqAllowed
		}
		maxFreqAllowed = max(0, freqArray[i]-1)
	}

	return deleteCount

}

func TestMinDeletions() {
	fmt.Println(minDeletions("aabbcccdddef"))
}
