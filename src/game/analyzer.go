package game

func checkWin() bool {
	return checkRow() || checkColumn() || checkDiagonal()
}

func checkRow() bool {
	for i := 0; i < 9; i += 3 {
		if board[i] == ' ' {
			continue
		}
		if (board[i] == board[i+1]) && (board[i+1] == board[i+2]) {
			return true
		}
	}
	return false
}

func checkColumn() bool {
	for i := 0; i < 2; i++ {
		if board[i] == ' ' {
			continue
		}
		if (board[i] == board[i+3]) && (board[i+3] == board[i+6]) {
			return true
		}
	}
	return false
}

func checkDiagonal() bool {
	if board[4] == ' ' {
		return false
	}
	if (board[0] == board[4]) && (board[4] == board[8]) {
		return true
	}

	if (board[2] == board[4]) && (board[4] == board[6]) {
		return true
	}

	return false
}
