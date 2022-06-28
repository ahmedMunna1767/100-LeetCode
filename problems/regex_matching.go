package problems

import "fmt"

const (
	NOTVISITED = iota
	TRUE
	FALSE
)

// Problem Two
// * Recursive Solution * //
func IsMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	// check if first character matches
	firstMatch := (!(s == "") && (p[0] == s[0] || p[0] == '.'))

	if len(p) >= 2 && p[1] == '*' {
		// if there is a klene star
		return IsMatch(s, p[2:]) || (firstMatch && IsMatch(s[1:], p))
	} else {
		// If 2nd character is not a klene start
		return firstMatch && (IsMatch(s[1:], p[1:]))
	}
}

// * DP solution
func IsMatchDynamic(s string, p string) bool {

	// Init the memoization array
	memo := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		memo[i] = make([]int, len(p)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(p)+1; j++ {
			memo[i][j] = NOTVISITED
		}
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// if the node is previously computed, return appropriate value
		if memo[i][j] != NOTVISITED {
			return TRUE == memo[i][j]
		}

		var ans bool
		if j == len(p) {
			// reach both end at the same time
			ans = (i == len(s))
		} else {
			// if the first characters match
			firstMatch := (i < len(s) && (p[j] == s[i] || p[j] == '.'))

			// in case of a kleene star
			if j+1 < len(p) && p[j+1] == '*' {
				ans = (dp(i, j+2) || firstMatch && dp(i+1, j))
			} else {
				ans = firstMatch && dp(i+1, j+1)
			}
		}
		if ans {
			memo[i][j] = TRUE
		} else {
			memo[i][j] = FALSE
		}
		return ans
	}

	return dp(0, 0)
}

func TestRegExpMatching() {
	fmt.Println(IsMatchDynamic("aa", ".."))
	fmt.Println(IsMatch("aa", ".."))
}
