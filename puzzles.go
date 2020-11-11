package main

import "./sudoku"

func thermoAMS() sudoku.Board {

	b := `
	    1  3 
	      8 4
	       1
	        5
	
	8
	 8
	2 3
	 4  2
	`
	board := sudoku.ParseBoard(b)

	board.AddClue(sudoku.NewThermometer(
		sudoku.NewCoordinate(4, 1),
		sudoku.NewCoordinate(3, 1),
		sudoku.NewCoordinate(2, 1),
		sudoku.NewCoordinate(1, 2),
		sudoku.NewCoordinate(2, 3),
		sudoku.NewCoordinate(3, 3),
		sudoku.NewCoordinate(4, 3),
	))

	board.AddClue(sudoku.NewThermometer(
		sudoku.NewCoordinate(3, 2),
		sudoku.NewCoordinate(3, 3),
		sudoku.NewCoordinate(4, 3),
	))

	board.AddClue(sudoku.NewThermometer(
		sudoku.NewCoordinate(6, 4),
		sudoku.NewCoordinate(5, 4),
		sudoku.NewCoordinate(4, 4),
		sudoku.NewCoordinate(5, 5),
		sudoku.NewCoordinate(4, 6),
		sudoku.NewCoordinate(5, 6),
		sudoku.NewCoordinate(6, 6),
	))

	board.AddClue(sudoku.NewThermometer(
		sudoku.NewCoordinate(7, 9),
		sudoku.NewCoordinate(7, 8),
		sudoku.NewCoordinate(7, 7),
		sudoku.NewCoordinate(8, 7),
		sudoku.NewCoordinate(8, 8),
		sudoku.NewCoordinate(8, 9),
		sudoku.NewCoordinate(9, 9),
		sudoku.NewCoordinate(9, 8),
		sudoku.NewCoordinate(9, 7),
	))

	return board
}
