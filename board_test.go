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
	possibleLocations := getPossibleMoves(location, ROOK)
	if len(possibleLocations) != 14 {
		t.Error("wrong number of 00 Rook possible moves")
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
	possibleLocations := getPossibleMoves(location, ROOK)
	if len(possibleLocations) != 14 {
		t.Error("wrong number of 07 Rook possible moves")
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
	possibleLocations := getPossibleMoves(location, ROOK)
	if len(possibleLocations) != 14 {
		t.Error("possible moves for 44 Rook are not 14")
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
	possibleLocations := getPossibleMoves(location, KNIGHT)
	if len(possibleLocations) != 3 {
		t.Error("possible moves for 01 Knight are not 3")
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
	possibleLocations := getPossibleMoves(location, KNIGHT)
	if len(possibleLocations) != 3 {
		t.Error("possible moves for 71 Knight are not 3")
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
	possibleLocations := getPossibleMoves(location, KNIGHT)
	if len(possibleLocations) != 8 {
		t.Error("possible moves for 43 Knight are not 8")
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
