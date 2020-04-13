package main

import (
	"testing"
)

func findLocation(t *testing.T, possibleLocations []Location, location Location) {
	for _, l := range possibleLocations {
		if l.row == location.row && l.col == location.col {
			return
		}
	}
	t.Errorf("failure, possible location %d:%d not found\n", location.row, location.col)
}

func TestGetPossibleMovesRook00(t *testing.T) {
	location := Location{
		row: 0,
		col: 0,
	}
	possibleLocations := ROOK.GetPossibleMoves(location)
	if len(possibleLocations) != 14 {
		t.Errorf("possible moves for 00 Rook were expected to be 14 but they are %d", len(possibleLocations))
	}
	findLocation(t, possibleLocations, Location{row: 1, col: 0})
	findLocation(t, possibleLocations, Location{row: 2, col: 0})
	findLocation(t, possibleLocations, Location{row: 3, col: 0})
	findLocation(t, possibleLocations, Location{row: 4, col: 0})
	findLocation(t, possibleLocations, Location{row: 5, col: 0})
	findLocation(t, possibleLocations, Location{row: 6, col: 0})
	findLocation(t, possibleLocations, Location{row: 7, col: 0})
	findLocation(t, possibleLocations, Location{row: 0, col: 1})
	findLocation(t, possibleLocations, Location{row: 0, col: 2})
	findLocation(t, possibleLocations, Location{row: 0, col: 3})
	findLocation(t, possibleLocations, Location{row: 0, col: 4})
	findLocation(t, possibleLocations, Location{row: 0, col: 5})
	findLocation(t, possibleLocations, Location{row: 0, col: 6})
	findLocation(t, possibleLocations, Location{row: 0, col: 7})
}

func TestGetPossibleMovesRook07(t *testing.T) {
	location := Location{
		row: 0,
		col: 7,
	}
	possibleLocations := ROOK.GetPossibleMoves(location)
	if len(possibleLocations) != 14 {
		t.Errorf("possible moves for 07 Rook were expected to be 14 but they are %d", len(possibleLocations))

	}
	findLocation(t, possibleLocations, Location{row: 1, col: 7})
	findLocation(t, possibleLocations, Location{row: 2, col: 7})
	findLocation(t, possibleLocations, Location{row: 3, col: 7})
	findLocation(t, possibleLocations, Location{row: 4, col: 7})
	findLocation(t, possibleLocations, Location{row: 5, col: 7})
	findLocation(t, possibleLocations, Location{row: 6, col: 7})
	findLocation(t, possibleLocations, Location{row: 7, col: 7})
	findLocation(t, possibleLocations, Location{row: 0, col: 0})
	findLocation(t, possibleLocations, Location{row: 0, col: 1})
	findLocation(t, possibleLocations, Location{row: 0, col: 2})
	findLocation(t, possibleLocations, Location{row: 0, col: 3})
	findLocation(t, possibleLocations, Location{row: 0, col: 4})
	findLocation(t, possibleLocations, Location{row: 0, col: 5})
	findLocation(t, possibleLocations, Location{row: 0, col: 6})
}

func TestGetPossibleMovesRook44(t *testing.T) {
	location := Location{
		row: 4,
		col: 4,
	}
	possibleLocations := ROOK.GetPossibleMoves(location)
	if len(possibleLocations) != 14 {
		t.Errorf("possible moves for 44 Rook were expected to be 14 but they are %d", len(possibleLocations))
	}
	findLocation(t, possibleLocations, Location{row: 0, col: 4})
	findLocation(t, possibleLocations, Location{row: 1, col: 4})
	findLocation(t, possibleLocations, Location{row: 2, col: 4})
	findLocation(t, possibleLocations, Location{row: 3, col: 4})
	findLocation(t, possibleLocations, Location{row: 5, col: 4})
	findLocation(t, possibleLocations, Location{row: 6, col: 4})
	findLocation(t, possibleLocations, Location{row: 7, col: 4})
	findLocation(t, possibleLocations, Location{row: 4, col: 0})
	findLocation(t, possibleLocations, Location{row: 4, col: 1})
	findLocation(t, possibleLocations, Location{row: 4, col: 2})
	findLocation(t, possibleLocations, Location{row: 4, col: 3})
	findLocation(t, possibleLocations, Location{row: 4, col: 5})
	findLocation(t, possibleLocations, Location{row: 4, col: 6})
	findLocation(t, possibleLocations, Location{row: 4, col: 7})
}

func TestGetPossibleMovesKnight01(t *testing.T) {
	location := Location{
		row: 0,
		col: 1,
	}
	possibleLocations := KNIGHT.GetPossibleMoves(location)
	if len(possibleLocations) != 3 {
		t.Errorf("possible moves for 01 Knight were expected to be 3 but they are %d", len(possibleLocations))
	}
	findLocation(t, possibleLocations, Location{row: 2, col: 0})
	findLocation(t, possibleLocations, Location{row: 2, col: 2})
	findLocation(t, possibleLocations, Location{row: 1, col: 3})
}

func TestGetPossibleMovesKnight71(t *testing.T) {
	location := Location{
		row: 7,
		col: 1,
	}
	possibleLocations := KNIGHT.GetPossibleMoves(location)
	if len(possibleLocations) != 3 {
		t.Errorf("possible moves for 71 Knight were expected to be 3 but they are %d", len(possibleLocations))
	}
	findLocation(t, possibleLocations, Location{row: 5, col: 0})
	findLocation(t, possibleLocations, Location{row: 5, col: 2})
	findLocation(t, possibleLocations, Location{row: 6, col: 3})
}

func TestGetPossibleMovesKnight43(t *testing.T) {
	location := Location{
		row: 4,
		col: 3,
	}
	possibleLocations := KNIGHT.GetPossibleMoves(location)
	if len(possibleLocations) != 8 {
		t.Errorf("possible moves for 43 Knight were expected to be 8 but they are %d", len(possibleLocations))
	}

	// top hand
	findLocation(t, possibleLocations, Location{row: 2, col: 2})
	findLocation(t, possibleLocations, Location{row: 2, col: 4})

	// right hand
	findLocation(t, possibleLocations, Location{row: 3, col: 5})
	findLocation(t, possibleLocations, Location{row: 5, col: 5})

	// bottom hand
	findLocation(t, possibleLocations, Location{row: 6, col: 2})
	findLocation(t, possibleLocations, Location{row: 6, col: 4})

	// left hand
	findLocation(t, possibleLocations, Location{row: 5, col: 1})
	findLocation(t, possibleLocations, Location{row: 3, col: 1})
}

func TestGetPossibleMovesBishop02(t *testing.T) {
	location := Location{
		row: 0,
		col: 2,
	}
	possibleLocations := BISHOP.GetPossibleMoves(location)
	if len(possibleLocations) != 7 {
		t.Errorf("possible moves for 02 Bishop were expected to be 7 but they are %d", len(possibleLocations))
	}

	// bottom-left
	findLocation(t, possibleLocations, Location{row: 2, col: 0})
	findLocation(t, possibleLocations, Location{row: 1, col: 1})

	// bottom-right
	findLocation(t, possibleLocations, Location{row: 1, col: 3})
	findLocation(t, possibleLocations, Location{row: 2, col: 4})
	findLocation(t, possibleLocations, Location{row: 3, col: 5})
	findLocation(t, possibleLocations, Location{row: 4, col: 6})
	findLocation(t, possibleLocations, Location{row: 5, col: 7})
}

func TestGetPossibleMovesBishop33(t *testing.T) {
	location := Location{
		row: 3,
		col: 3,
	}
	possibleLocations := BISHOP.GetPossibleMoves(location)
	if len(possibleLocations) != 13 {
		t.Errorf("possible moves for 33 Bishop were expected to be 13 but they are %d", len(possibleLocations))
	}

	// top-right
	findLocation(t, possibleLocations, Location{row: 2, col: 4})
	findLocation(t, possibleLocations, Location{row: 1, col: 5})
	findLocation(t, possibleLocations, Location{row: 0, col: 6})

	// bottom-right
	findLocation(t, possibleLocations, Location{row: 4, col: 4})
	findLocation(t, possibleLocations, Location{row: 5, col: 5})
	findLocation(t, possibleLocations, Location{row: 6, col: 6})
	findLocation(t, possibleLocations, Location{row: 7, col: 7})

	// bottom-left
	findLocation(t, possibleLocations, Location{row: 4, col: 2})
	findLocation(t, possibleLocations, Location{row: 5, col: 1})
	findLocation(t, possibleLocations, Location{row: 6, col: 0})

	// top-left
	findLocation(t, possibleLocations, Location{row: 0, col: 0})
	findLocation(t, possibleLocations, Location{row: 1, col: 1})
	findLocation(t, possibleLocations, Location{row: 2, col: 2})
}

func TestGetPossibleMovesQueen03(t *testing.T) {
	location := Location{
		row: 0,
		col: 3,
	}
	possibleLocations := QUEEN.GetPossibleMoves(location)
	if len(possibleLocations) != 21 {
		t.Errorf("possible moves for 03 Queen were expected to be 21 but they were %d", len(possibleLocations))
	}

	// horizontal left
	findLocation(t, possibleLocations, Location{row: 0, col: 2})
	findLocation(t, possibleLocations, Location{row: 0, col: 1})
	findLocation(t, possibleLocations, Location{row: 0, col: 0})

	// horizontal right
	findLocation(t, possibleLocations, Location{row: 0, col: 4})
	findLocation(t, possibleLocations, Location{row: 0, col: 5})
	findLocation(t, possibleLocations, Location{row: 0, col: 6})
	findLocation(t, possibleLocations, Location{row: 0, col: 7})

	// diagonal bottom-left
	findLocation(t, possibleLocations, Location{row: 1, col: 2})
	findLocation(t, possibleLocations, Location{row: 2, col: 1})
	findLocation(t, possibleLocations, Location{row: 3, col: 0})

	// diagonal bottom-right
	findLocation(t, possibleLocations, Location{row: 1, col: 4})
	findLocation(t, possibleLocations, Location{row: 2, col: 5})
	findLocation(t, possibleLocations, Location{row: 3, col: 6})
	findLocation(t, possibleLocations, Location{row: 4, col: 7})
}

func TestGetPossibleMovesQueen43(t *testing.T) {
	location := Location{
		row: 4,
		col: 3,
	}
	possibleLocations := QUEEN.GetPossibleMoves(location)
	if len(possibleLocations) != 27 {
		t.Errorf("possible moves for 43 Queen were expected to be 27 but they are %d", len(possibleLocations))
	}

	// vertical top
	findLocation(t, possibleLocations, Location{row: 3, col: 3})
	findLocation(t, possibleLocations, Location{row: 2, col: 3})
	findLocation(t, possibleLocations, Location{row: 1, col: 3})
	findLocation(t, possibleLocations, Location{row: 0, col: 3})

	// vertical bottom
	findLocation(t, possibleLocations, Location{row: 5, col: 3})
	findLocation(t, possibleLocations, Location{row: 6, col: 3})
	findLocation(t, possibleLocations, Location{row: 7, col: 3})

	// horizontal left
	findLocation(t, possibleLocations, Location{row: 4, col: 2})
	findLocation(t, possibleLocations, Location{row: 4, col: 1})
	findLocation(t, possibleLocations, Location{row: 4, col: 0})

	// horizontal right
	findLocation(t, possibleLocations, Location{row: 4, col: 4})
	findLocation(t, possibleLocations, Location{row: 4, col: 5})
	findLocation(t, possibleLocations, Location{row: 4, col: 6})
	findLocation(t, possibleLocations, Location{row: 4, col: 7})

	// diagonal top-right
	findLocation(t, possibleLocations, Location{row: 3, col: 4})
	findLocation(t, possibleLocations, Location{row: 2, col: 5})
	findLocation(t, possibleLocations, Location{row: 1, col: 6})
	findLocation(t, possibleLocations, Location{row: 0, col: 7})

	// diagonal bottom-right
	findLocation(t, possibleLocations, Location{row: 5, col: 4})
	findLocation(t, possibleLocations, Location{row: 6, col: 5})
	findLocation(t, possibleLocations, Location{row: 7, col: 6})

	// diagonal bottom-left
	findLocation(t, possibleLocations, Location{row: 5, col: 2})
	findLocation(t, possibleLocations, Location{row: 6, col: 1})
	findLocation(t, possibleLocations, Location{row: 7, col: 0})

	// diagonal top-left
	findLocation(t, possibleLocations, Location{row: 3, col: 2})
	findLocation(t, possibleLocations, Location{row: 2, col: 1})
	findLocation(t, possibleLocations, Location{row: 1, col: 0})
}

func TestGetPossibleMovesKing04(t *testing.T) {
	location := Location{
		row: 0,
		col: 4,
	}
	possibleLocations := KING.GetPossibleMoves(location)
	if len(possibleLocations) != 5 {
		t.Errorf("possible moves for 04 King were expected to be 5 but they were %d", len(possibleLocations))
	}

	// counterclockwise
	findLocation(t, possibleLocations, Location{row: 0, col: 3})
	findLocation(t, possibleLocations, Location{row: 1, col: 3})
	findLocation(t, possibleLocations, Location{row: 1, col: 4})
	findLocation(t, possibleLocations, Location{row: 1, col: 5})
	findLocation(t, possibleLocations, Location{row: 0, col: 5})
}

func TestGetPossibleMovesKing35(t *testing.T) {
	location := Location{
		row: 3,
		col: 5,
	}
	possibleLocations := KING.GetPossibleMoves(location)
	if len(possibleLocations) != 8 {
		t.Errorf("possible moves for 35 King were expected to be 8 but they were %d", len(possibleLocations))
	}

	// clockwise
	findLocation(t, possibleLocations, Location{row: 2, col: 5})
	findLocation(t, possibleLocations, Location{row: 2, col: 6})
	findLocation(t, possibleLocations, Location{row: 3, col: 6})
	findLocation(t, possibleLocations, Location{row: 4, col: 6})
	findLocation(t, possibleLocations, Location{row: 4, col: 5})
	findLocation(t, possibleLocations, Location{row: 4, col: 4})
	findLocation(t, possibleLocations, Location{row: 3, col: 4})
	findLocation(t, possibleLocations, Location{row: 2, col: 4})
}

func TestGetPossibleMovesPawn65(t *testing.T) {
	location := Location{
		row: 6,
		col: 5,
	}
	possibleLocations := PAWN.GetPossibleMoves(location)
	if len(possibleLocations) != 7 {
		t.Errorf("possible moves for 65 Pawn were expected to be 7 but they were %d", len(possibleLocations))
	}

	// forward once
	findLocation(t, possibleLocations, Location{row: 5, col: 5})
	// forward twice
	findLocation(t, possibleLocations, Location{row: 4, col: 5})
	// forward left capture
	findLocation(t, possibleLocations, Location{row: 5, col: 4})
	// forward right capture
	findLocation(t, possibleLocations, Location{row: 5, col: 6})
	// backwards once
	findLocation(t, possibleLocations, Location{row: 7, col: 5})
	// backwards capture left
	findLocation(t, possibleLocations, Location{row: 7, col: 4})
	// backwards capture right
	findLocation(t, possibleLocations, Location{row: 7, col: 6})
}

func TestGetPossibleMovesPawn22(t *testing.T) {
	location := Location{
		row: 2,
		col: 2,
	}
	possibleLocations := PAWN.GetPossibleMoves(location)
	if len(possibleLocations) != 6 {
		t.Errorf("possible moves for 22 Pawn were expected to be 6 but they were %d", len(possibleLocations))
	}

	// forward once
	findLocation(t, possibleLocations, Location{row: 3, col: 2})
	// forward left capture
	findLocation(t, possibleLocations, Location{row: 3, col: 1})
	// forward right capture
	findLocation(t, possibleLocations, Location{row: 3, col: 3})
	// backwards once
	findLocation(t, possibleLocations, Location{row: 1, col: 2})
	// backwards capture left
	findLocation(t, possibleLocations, Location{row: 1, col: 1})
	// backwards capture right
	findLocation(t, possibleLocations, Location{row: 1, col: 3})
}
