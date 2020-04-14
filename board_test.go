package main

import (
	"testing"
)

func TestBoardPawn(t *testing.T) {
	board := Board{}
	board.Init()

	// create move
	turn := WHITE
	command := "e7 e5"
	move, _ := NewMove(board, turn, command)

	// execute move
	board.Execute(move)

	// verify piece moved
	locationBefore := move.GetLocation(BEFORE)
	locationAfter := move.GetLocation(AFTER)
	if board[locationBefore.row][locationBefore.col] != "   " {
		t.Error("white pawn did not move out of e7")
	}
	if board[locationAfter.row][locationAfter.col] != "○ P" {
		t.Error("white pawn did not move into e5")
	}
}
