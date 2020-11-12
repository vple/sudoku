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

func killerTest() sudoku.Board {
	b := `` // Empty
	board := sudoku.ParseBoard(b)

	board.AddClue(sudoku.NewKillerCage(
		[]sudoku.Coordinate{
			sudoku.NewCoordinate(1, 1),
			sudoku.NewCoordinate(1, 2),
			sudoku.NewCoordinate(1, 3),
		},
		6,
	))

	return board
}

func ctcCompetition() sudoku.Board {
	b := `` // Empty
	board := sudoku.ParseBoard(b)

	board.AddClue(sudoku.NewKillerCage(
		[]sudoku.Coordinate{
			sudoku.NewCoordinate(3, 4),
			sudoku.NewCoordinate(3, 5),
			sudoku.NewCoordinate(4, 4),
			sudoku.NewCoordinate(4, 5),
		},
		25,
	))
	board.AddClue(sudoku.NewKillerCage(
		[]sudoku.Coordinate{
			sudoku.NewCoordinate(4, 6),
			sudoku.NewCoordinate(4, 7),
			sudoku.NewCoordinate(5, 6),
			sudoku.NewCoordinate(5, 7),
		},
		27,
	))
	board.AddClue(sudoku.NewKillerCage(
		[]sudoku.Coordinate{
			sudoku.NewCoordinate(5, 3),
			sudoku.NewCoordinate(5, 4),
			sudoku.NewCoordinate(6, 3),
			sudoku.NewCoordinate(6, 4),
		},
		29,
	))
	board.AddClue(sudoku.NewKillerCage(
		[]sudoku.Coordinate{
			sudoku.NewCoordinate(6, 5),
			sudoku.NewCoordinate(6, 6),
			sudoku.NewCoordinate(7, 5),
			sudoku.NewCoordinate(7, 6),
		},
		28,
	))

	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 2), sudoku.NewCoordinate(2, 1)),
		5,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(3, 1), sudoku.NewCoordinate(9, 7)),
		50,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(2, 9), sudoku.NewCoordinate(1, 8)),
		5,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(4, 9), sudoku.NewCoordinate(1, 6)),
		9,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(8, 1), sudoku.NewCoordinate(9, 2)),
		5,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(9, 8), sudoku.NewCoordinate(8, 9)),
		5,
	))

	return board
}

func littleKiller() sudoku.Board {
	// Actual puzzle has no given values.
	b := `
	  2149
	  6 7
	 5 8 217
	 7   1	
	      4
	3
	9 3

	62 3 4
	` // Empty
	board := sudoku.ParseBoard(b)

	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 4), sudoku.NewCoordinate(4, 1)),
		14,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 5), sudoku.NewCoordinate(5, 1)),
		33,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 6), sudoku.NewCoordinate(6, 1)),
		40,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 7), sudoku.NewCoordinate(7, 1)),
		46,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(6, 1), sudoku.NewCoordinate(9, 4)),
		18,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(9, 6), sudoku.NewCoordinate(6, 9)),
		16,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(5, 9), sudoku.NewCoordinate(1, 5)),
		13,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(6, 9), sudoku.NewCoordinate(1, 4)),
		24,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(7, 9), sudoku.NewCoordinate(1, 3)),
		24,
	))

	return board
}

func lightsaber() sudoku.Board {
	b := `` // Empty
	board := sudoku.ParseBoard(b)

	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(1, 3),
		sudoku.NewCoordinate(2, 3),
	))
	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(3, 1),
		sudoku.NewCoordinate(3, 2),
	))
	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(4, 5),
		sudoku.NewCoordinate(4, 6),
	))
	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(5, 4),
		sudoku.NewCoordinate(6, 4),
	))
	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(7, 8),
		sudoku.NewCoordinate(7, 9),
	))
	board.AddClue(sudoku.Sum5(
		sudoku.NewCoordinate(8, 7),
		sudoku.NewCoordinate(9, 7),
	))

	board.AddClue(sudoku.Sum10(
		sudoku.NewCoordinate(1, 9),
		sudoku.NewCoordinate(2, 9),
	))
	board.AddClue(sudoku.Sum10(
		sudoku.NewCoordinate(2, 6),
		sudoku.NewCoordinate(2, 7),
	))
	board.AddClue(sudoku.Sum10(
		sudoku.NewCoordinate(8, 3),
		sudoku.NewCoordinate(9, 3),
	))

	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 1), sudoku.NewCoordinate(9, 9)),
		57,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(1, 2), sudoku.NewCoordinate(8, 9)),
		31,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(2, 1), sudoku.NewCoordinate(9, 8)),
		32,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(3, 9), sudoku.NewCoordinate(9, 3)),
		32,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(4, 9), sudoku.NewCoordinate(9, 4)),
		25,
	))
	board.AddClue(sudoku.NewLittleKiller(
		board.Diagonal(sudoku.NewCoordinate(5, 9), sudoku.NewCoordinate(9, 5)),
		16,
	))

	return board
}
