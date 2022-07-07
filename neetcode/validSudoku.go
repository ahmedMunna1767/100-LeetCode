package neetcode

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		hash := make(map[byte]bool, 0)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := hash[board[i][j]]; ok {
				return false
			} else {
				hash[board[i][j]] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		hash := make(map[byte]bool, 0)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, ok := hash[board[j][i]]; ok {
				return false
			} else {
				hash[board[j][i]] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		hash := make(map[byte]bool, 0)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if board[j+i*3][k+i*3] == '.' {
					continue
				}
				if _, ok := hash[board[j+i*3][k+i*3]]; ok {
					return false
				} else {
					hash[board[k+i*3][j+i*3]] = true
				}
			}
		}
	}

	return true
}
