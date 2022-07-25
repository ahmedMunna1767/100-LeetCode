package contest

func bestHand(ranks []int, suits []byte) string {
	ranksMap := make(map[int]int)
	suitsMap := make(map[byte]int)

	for i := 0; i < len(suits); i++ {
		suitsMap[suits[i]]++
	}

	for i := 0; i < len(ranks); i++ {
		ranksMap[ranks[i]]++
	}

	for _, v := range suitsMap {
		if v == 5 {
			return "Flush"
		}
	}

	for _, v := range ranksMap {
		if v >= 3 {
			return "Three of a Kind"
		}
		if v == 2 {
			return "Pair"
		}
	}

	return "High Card"
}

func zeroFilledSubarray(nums []int) int64 {
	count := 0
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			zeroCount++
			zeroCount = zeroCount + count - 1
		}
	}

	return int64(zeroCount)
}
