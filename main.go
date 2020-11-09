package main

import (
	"fmt"
	"time"

	"./sat"
	"./sudoku"
	"./sudoku/conversion"
)

func easySudoku() sudoku.Board {
	b := `
	8 76 3   
	6 9  7831
	 31  46  
	     6 84
	21 7   9 
	4  8 215 
	     8  5
	  542   3
	3745  9 8
	`
	return sudoku.ParseBoard(b)
}

func thermometerSudoku() sudoku.Board {
	b := "" // Empty
	board := sudoku.ParseBoard(b)

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(1, 2),
		sudoku.NewCoordinate(1, 3),
		sudoku.NewCoordinate(1, 4),
		sudoku.NewCoordinate(2, 4),
		sudoku.NewCoordinate(2, 3),
		sudoku.NewCoordinate(2, 2),
		sudoku.NewCoordinate(3, 2),
		sudoku.NewCoordinate(3, 3),
		sudoku.NewCoordinate(3, 4),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(4, 4),
		sudoku.NewCoordinate(4, 3),
		sudoku.NewCoordinate(4, 2),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(5, 2),
		sudoku.NewCoordinate(6, 2),
		sudoku.NewCoordinate(7, 2),
		sudoku.NewCoordinate(8, 2),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(8, 3),
		sudoku.NewCoordinate(8, 4),
		sudoku.NewCoordinate(7, 4),
		sudoku.NewCoordinate(6, 4),
		sudoku.NewCoordinate(5, 4),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(2, 7),
		sudoku.NewCoordinate(2, 8),
		sudoku.NewCoordinate(2, 9),
		sudoku.NewCoordinate(3, 9),
		sudoku.NewCoordinate(3, 8),
		sudoku.NewCoordinate(3, 7),
		sudoku.NewCoordinate(4, 7),
		sudoku.NewCoordinate(4, 8),
		sudoku.NewCoordinate(4, 9),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(9, 6),
		sudoku.NewCoordinate(8, 6),
		sudoku.NewCoordinate(7, 6),
		sudoku.NewCoordinate(6, 6),
	))

	board.AddConstraint(sudoku.NewIncreasingValueConstraint(
		sudoku.NewCoordinate(6, 7),
		sudoku.NewCoordinate(6, 8),
		sudoku.NewCoordinate(7, 8),
		sudoku.NewCoordinate(8, 8),
		sudoku.NewCoordinate(9, 8),
		sudoku.NewCoordinate(9, 7),
	))

	return board
}

func diagonalSudoku() sudoku.Board {
	b := `
	 7 3 5 1 
	9  6 1  3
	3       4
	    8    
	   1 6   
	  3   4  
	   8 2   
	6       1
	  7   9  
	`

	board := sudoku.ParseBoard(b)
	// rules := sudoku.Rules{
	// 	diagonalsUnique: true,
	// }
	// board.rules = rules

	return board
}

func solverTest(board sudoku.Board) {
	fmt.Println("Input:")
	fmt.Println(board)

	formula := conversion.ToFormula(board)

	fmt.Println("\nSolving...")
	start := time.Now()
	results, ok := sat.Solve(formula, make(map[string]bool))
	duration := time.Since(start)

	board = conversion.ParseState(results)
	fmt.Println("\nResult:", ok, duration)
	fmt.Println(board)
}

func main() {
	solverTest(thermometerSudoku())
}
