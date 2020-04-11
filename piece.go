package main

// Piece defines all chess pieces of the board
type Piece int

const (
	// PAWN is the pawn chess piece of a chess board
	PAWN Piece = iota
	// ROOK is the rook chess piece of a chess board
	ROOK
	// KNIGHT is the knight chess piece of a chess board
	KNIGHT
	// BISHOP is the bishop chess piece of a chess board
	BISHOP
	// QUEEN is the queen chess piece of a chess board
	QUEEN
	// KING is the king chess piece of a chess board
	KING
)
