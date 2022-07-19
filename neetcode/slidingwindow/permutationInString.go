package slidingwindow

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1map := [26]int{}
	s2map := [26]int{}

	for i := 0; i < len(s1); i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}
	fmt.Println(s1map)
	fmt.Println(s2map)

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}
	fmt.Println(count)

	for i := 0; i < len(s2)-len(s1); i++ {
		r := s2[i+len(s1)] - 'a'
		l := s2[i] - 'a'
		if count == 26 {
			return true
		}

		s2map[r]++
		if s2map[r] == s1map[r] {
			count++
		} else if s2map[r] == s1map[r]+1 {
			count--
		}

		s2map[l]--
		if s2map[l] == s1map[l] {
			count++
		} else if s2map[l] == s1map[l]-1 {
			count--
		}
		fmt.Println(s2map, count)
	}
	return count == 26
}

func CheckInclusion() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
