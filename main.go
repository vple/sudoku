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
	solverTest(easySudoku())
}
