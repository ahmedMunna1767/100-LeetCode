package slidingwindow

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	// Dictionary for target string
	tMap := make(map[byte]int, 0)
	// Dictionary for Window string
	wMap := make(map[byte]int, 0)
	minLength := math.MaxInt
	have, left, right, ansLeft, required := 0, 0, 0, 0, 0

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	required = len(tMap) // length of target dictionary

	for right < len(s) {
		char := s[right]
		wMap[char]++

		if _, ok := tMap[char]; ok {
			if tMap[char] == wMap[char] {
				have++
			}
		}

		for have == required && left <= right {
			windowLength := right - left + 1
			if windowLength < minLength {
				minLength = windowLength
				ansLeft = left
			}
			char = s[left]
			left++
			wMap[char]--
			if _, ok := tMap[char]; ok {
				if wMap[char] < tMap[char] {
					have--
				}
			}
		}
		right++

	}

	if minLength == math.MaxInt {
		return ""
	}
	return s[ansLeft : ansLeft+minLength]
}

func M(s string, t string) string {
	have := 0
	left := 0
	right := len(s)
	found := false
	minLength := len(s)
	l, r := 0, 0

	tmap := make(map[byte]int, 0)
	smap := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		tmap[t[i]] += 1
	}
	need := len(tmap)

	for r = 0; r < len(s); r++ {
		smap[s[r]] += 1

		if _, ok := tmap[s[r]]; ok {
			if tmap[s[r]] == smap[s[r]] {
				have++
			}
		}
		// fmt.Printf("%v, %v, %d	\n", tmap, smap, have)
		for have == need && l <= r {
			found = true
			if r-l+1 < minLength {
				minLength = r - l + 1
				left = l
				right = r
			}
			smap[s[l]] -= 1
			if _, ok := tmap[s[l]]; ok {
				if smap[s[l]] < tmap[s[l]] {
					have--
				}
			}
			l++
		}
	}

	fmt.Println(left, right)
	if !found {
		return ""
	}
	return s[left : left+minLength]
}

func MinWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("AA", "AA"))

}
