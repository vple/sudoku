package main

import (
	"fmt"
	"time"
)

type cell struct {
	row            int
	col            int
	value          int
	possibleValues []int
}

func newCell(row, col int) *cell {
	c := new(cell)
	c.row = row
	c.col = col
	c.possibleValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	return c
}

func (c cell) variableName() string {
	return fmt.Sprintf("%d%d", c.row, c.col)
}

func (c cell) getValue() (int, bool) {
	return c.value, c.value != 0
}

func diagonalSudoku() Board {
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

	board := ParseBoard(b)
	rules := Rules{
		diagonalsUnique: true,
	}
	board.rules = rules

	return board
}

func sampleBoard() Board {
	initialValues := map[Coordinate]int{
		NewCoordinate(1, 3): 7,
		NewCoordinate(2, 4): 5,
		NewCoordinate(2, 5): 4,
		NewCoordinate(2, 8): 9,
		NewCoordinate(2, 9): 3,
		NewCoordinate(3, 1): 8,
		NewCoordinate(3, 6): 6,
		NewCoordinate(3, 7): 2,
		NewCoordinate(3, 8): 5,
		NewCoordinate(4, 9): 1,
		NewCoordinate(5, 2): 1,
		NewCoordinate(5, 4): 3,
		NewCoordinate(5, 7): 5,
		NewCoordinate(5, 9): 7,
		NewCoordinate(6, 8): 2,
		NewCoordinate(7, 2): 4,
		NewCoordinate(7, 5): 7,
		NewCoordinate(8, 1): 6,
		NewCoordinate(8, 7): 8,
		NewCoordinate(9, 3): 3,
		NewCoordinate(9, 4): 2,
		NewCoordinate(9, 6): 4,
	}

	return NewStandardBoard(initialValues)
}

func sampleBoard2() Board {
	initialValues := map[Coordinate]int{
		NewCoordinate(1, 3): 6,
		NewCoordinate(1, 6): 1,
		NewCoordinate(1, 8): 2,
		NewCoordinate(1, 9): 8,
		NewCoordinate(2, 2): 9,
		NewCoordinate(2, 3): 8,
		NewCoordinate(2, 4): 3,
		NewCoordinate(2, 9): 5,
		NewCoordinate(3, 4): 5,
		NewCoordinate(4, 5): 7,
		NewCoordinate(4, 6): 5,
		NewCoordinate(4, 8): 1,
		NewCoordinate(4, 9): 2,
		NewCoordinate(5, 2): 5,
		NewCoordinate(5, 8): 7,
		NewCoordinate(6, 1): 8,
		NewCoordinate(6, 2): 7,
		NewCoordinate(6, 4): 4,
		NewCoordinate(6, 5): 3,
		NewCoordinate(7, 6): 4,
		NewCoordinate(8, 1): 7,
		NewCoordinate(8, 6): 9,
		NewCoordinate(8, 7): 5,
		NewCoordinate(8, 8): 8,
		NewCoordinate(9, 1): 6,
		NewCoordinate(9, 2): 8,
		NewCoordinate(9, 4): 2,
		NewCoordinate(9, 7): 1,
	}

	return NewStandardBoard(initialValues)
}

func solverTest(board Board) {
	fmt.Println("Input:")
	fmt.Println(board)

	formula := board.Formula()

	fmt.Println("\nSolving...")
	start := time.Now()
	results, ok := Solve(formula, make(map[string]bool))
	duration := time.Since(start)

	board = ParseState(results)
	fmt.Println("\nResult:", ok, duration)
	fmt.Println(board)
}

func main() {
	solverTest(diagonalSudoku())
}
