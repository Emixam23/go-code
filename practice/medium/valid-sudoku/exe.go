package valid_sudoku

func checkX(board [][]byte) bool {
	for x := 0; x < 9; x++ {
		var container map[byte]int = make(map[byte]int)
		for y := 0; y < 9; y++ {
			cr := board[x][y]
			if cr == '.' {
				continue
			} else if _, ok := container[cr]; ok {
				return false
			}
			container[cr]++
		}
	}

	return true
}

func checkY(board [][]byte) bool {
	for y := 0; y < 9; y++ {
		var container map[byte]int = make(map[byte]int)
		for x := 0; x < 9; x++ {
			cr := board[x][y]
			if cr == '.' {
				continue
			} else if _, ok := container[cr]; ok {
				return false
			}
			container[cr]++
		}
	}

	return true
}

const bcSize = 3

func checkBCs(board [][]byte) bool {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if !checkBC(board, x*3, y*3) {
				return false
			}
		}
	}
	return true
}

func checkBC(board [][]byte, i, j int) bool {
	var container map[byte]int = make(map[byte]int)
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			cr := board[x+i][y+j]
			if cr == '.' {
				continue
			} else if _, ok := container[cr]; ok {
				return false
			}
			container[cr]++
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	if !checkX(board) {
		return false
	} else if !checkY(board) {
		return false
	} else if !checkBCs(board) {
		return false
	}

	return true
}
