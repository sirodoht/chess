package main

import (
	"testing"
)

func findLocation(t *testing.T, piece Piece, possibleLocations []Location, location Location) {
	for _, l := range possibleLocations {
		if l.row == location.row && l.col == location.col {
			return
		}
	}
	t.Errorf("failed for %s, possible location %d:%d not found\n", GetPieceName(ROOK), location.row, location.col)
}

func TestGetPossibleMovesRook(t *testing.T) {
	location := Location{
		row: 0,
		col: 0,
	}
	possibleLocations := getPossibleMoves(location, ROOK)
	if len(possibleLocations) != 14 {
		t.Error("wrong number of Rook possible moves")
	}
	findLocation(t, ROOK, possibleLocations, Location{row: 1, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 2, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 3, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 4, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 5, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 6, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 7, col: 0})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 1})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 2})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 3})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 4})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 5})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 6})
	findLocation(t, ROOK, possibleLocations, Location{row: 0, col: 7})
}
