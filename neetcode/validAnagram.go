package neetcode

import "reflect"

func ValidAnagram(s string, t string) bool {
	smap := make(map[byte]int)
	tmap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := smap[s[i]]; ok {
			smap[s[i]] = smap[s[i]] + 1
		} else {
			smap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := tmap[t[i]]; ok {
			tmap[s[i]] = tmap[t[i]] + 1
		} else {
			tmap[t[i]] = 1
		}
	}

	return reflect.DeepEqual(smap, tmap)
}
