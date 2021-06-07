package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			// 3x3

			for i2 := i - (i % 3); i2 < i+3-(i%3); i2++ {
				for j2 := j - (j % 3); j2 < j+3-(j%3); j2++ {
					if i == i2 && j == j2 {
						continue
					}
					if board[i][j] == board[i2][j2] {
						return false
					}
				}
			}

			// horizontal
			for k := j + 1; k < 9; k++ {
				if board[i][j] == board[i][k] {
					return false
				}
			}

			// vertical
			for k := i + 1; k < 9; k++ {
				if board[i][j] == board[k][j] {
					return false
				}
			}
		}
	}

	return true
}
