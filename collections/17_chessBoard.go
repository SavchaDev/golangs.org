package collections

import "fmt"

// ChessBoard расстановка фигур
func ChessBoard() {
	var board [8][8]rune // Массив из восьми массивов с восемью строками
	// black
	board[0] = [8]rune{9821, 9822, 9821, 9818, 9819, 9821, 9822, 9821}
	board[1] = [8]rune{9823, 9823, 9823, 9823, 9823, 9823, 9823, 9823}
	// white
	board[6] = [8]rune{9817, 9817, 9817, 9817, 9817, 9817, 9817, 9817}
	board[7] = [8]rune{9815, 9816, 9815, 9813, 9812, 9815, 9816, 9815}

	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Print("\n")
	}
}
