package picine

import "github.com/01-edu/z01"

/*
	Write a function that prints the solutions to the eight queens puzzle.

	Recursivity must be used to solve this problem.

	It should print something like this :

	$ go run .
	15863724
	16837425
	17468253
	...

	Each solution will be on a single line. The index of the placement of a queen starts at 1.

	It reads from left to right and each digit is the position for each column.

	The solutions will be printed in ascending order.
*/

func EightQueens() {
	var board [8]int
	solvQueens(board, 0)
}

func solvQueens(board [8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}
	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solvQueens(board, col+1)
		}
	}
}

func isSafe(board [8]int, row int, col int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]
		if prevRow == row {
			return false
		}
		if prevRow-row == prevCol-col || prevRow-row == col-prevCol {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '1'))
	}
	z01.PrintRune('\n')
}
