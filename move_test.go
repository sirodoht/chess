package main

import (
	"testing"
)

func TestRookMove(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "○ P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "h8 h6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Rook move not valid")
	}
}

func TestRookMoveInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "○ P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "h8 h4"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Rook move is valid")
	}
}

func TestRookMoveInvalidDiagonal(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "○ P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ R"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "   "},
	}

	// create move
	turn := WHITE
	command := "h7 g6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Rook move is valid")
	}
}

func TestRookCaptureTop(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"   ", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"● P", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "h8 h2"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Rook capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Rook capture strategy was not identified")
	}
}

func TestRookCaptureLeft(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"   ", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"● P", "   ", "   ", "   ", "   ", "   ", "   ", "○ R"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "   "},
	}

	// create move
	turn := WHITE
	command := "h3 a3"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Rook capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Rook capture strategy was not identified")
	}
}

func TestRookCaptureBottom(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"   ", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := BLACK
	command := "a1 a7"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Rook capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Rook capture strategy was not identified")
	}
}

func TestRookCaptureRight(t *testing.T) {
	board := Board{
		{"   ", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"   ", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"● R", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := BLACK
	command := "a4 g4"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Rook capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Rook capture strategy was not identified")
	}
}

func TestRookCaptureBlocked(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"   ", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"● P", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := BLACK
	command := "h1 h8"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Rook capture move is valid when it should not have been")
	}
}

func TestKnightMove(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "g8 h6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Knight move not valid")
	}
}

func TestKnightMoveInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "g8 h5"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Knight move is valid")
	}
}

func TestKnightCaptureTopRight(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "● P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ K", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "   ", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f6 g4"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Knight capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Knight capture strategy was not identified")
	}
}

func TestKnightCaptureTopLeft(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "   ", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "● P", "   ", "● P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ K", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "   ", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f6 e4"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Knight capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Knight capture strategy was not identified")
	}
}

func TestKnightCaptureLeftTop(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "   ", "● P", "● P", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "● P", "   "},
		{"   ", "   ", "   ", "● P", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ K", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "   ", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f6 d5"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Knight capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Knight capture strategy was not identified")
	}
}

func TestKnightCaptureBottomLeft(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "   ", "● P", "   ", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "○ P", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "● P", "   "},
		{"   ", "   ", "   ", "● P", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ K", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   ", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "   ", "○ R"},
	}

	// create move
	turn := BLACK
	command := "g1 f3"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Knight capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Knight capture strategy was not identified")
	}
}

func TestKnightCaptureInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "   ", "● P", "   ", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "● P", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "● P", "   "},
		{"   ", "   ", "   ", "● P", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ K", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   ", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "   ", "○ R"},
	}

	// create move
	turn := BLACK
	command := "g1 f3"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Knight capture move is valid when it should not have")
	}
}

func TestBishopMove(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f8 d6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Bishop move not valid")
	}
}

func TestBishopMoveInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "○ P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "   "},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f8 d6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Bishop move is valid")
	}
}

func TestBishopMoveInvalidVertical(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "○ P", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f8 f6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Bishop move is valid")
	}
}

func TestQueenMoveDiagonal(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "d8 h4"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Queen move not valid")
	}
}

func TestQueenMoveVertical(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "○ P", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "d8 d6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Queen move not valid")
	}
}

func TestQueenMoveInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "f8 f6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Queen move is valid")
	}
}

func TestKingMoveVertical(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e8 e7"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("King move not valid")
	}
}

func TestKingMoveDiagonal(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ G", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "   ", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e7 d6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("King move not valid")
	}
}

func TestKingMoveInvalidVertical(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e8 e6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("King move is valid")
	}
}

func TestKingMoveInvalidDiagonal(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ G", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "   ", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e7 c5"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("King move is valid")
	}
}

func TestPawnMove(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "h7 h6"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Pawn move not valid")
	}
}

func TestPawnMoveDouble(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "a7 a5"
	_, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("Pawn move not valid")
	}
}

func TestPawnMoveInvalid(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "h7 h4"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Pawn move is valid")
	}
}

func TestPawnMoveInvalidBackwards(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "   ", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "● P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e5 e6"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Pawn move is valid")
	}
}

func TestPawnMoveInvalidCapture(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "   ", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "● P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "○ P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "e5 e4"
	_, isValid, _ := NewMove(board, turn, command)
	if isValid {
		t.Error("Pawn move is valid")
	}
}

func TestPawnCaptureAsWhite(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "   ", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "● P", "   ", "   ", "   "},
		{"   ", "   ", "   ", "○ P", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "   ", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := WHITE
	command := "d5 e4"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Pawn capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Pawn capture strategy was not identified")
	}
}

func TestPawnCaptureAsBlack(t *testing.T) {
	board := Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "○ P", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}

	// create move
	turn := BLACK
	command := "h4 g5"
	move, isValid, _ := NewMove(board, turn, command)
	if !isValid {
		t.Error("invalid Pawn capture move")
	}
	if move.strategy != CAPTURE {
		t.Error("Pawn capture strategy was not identified")
	}
}
