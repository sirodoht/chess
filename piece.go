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

// GetPiece returns the Piece type given a piece notation
// e.g. "P" -> PAWN
func GetPiece(pieceNotation rune) Piece {
	pieces := map[rune]Piece{
		'P': PAWN,
		'R': ROOK,
		'K': KNIGHT,
		'B': BISHOP,
		'Q': QUEEN,
		'G': KING,
	}
	return pieces[pieceNotation]
}

// GetPieceName returns the name of the piece
// e.g. PAWN -> "Pawn"
func GetPieceName(piece Piece) string {
	pieceNames := map[Piece]string{
		PAWN:   "Pawn",
		ROOK:   "Rook",
		KNIGHT: "Knight",
		BISHOP: "Bishop",
		QUEEN:  "Queen",
		KING:   "King",
	}
	return pieceNames[piece]
}

// GetPieceAbbr returns the abbreviation of the piece
// e.g. PAWN -> "P"
func GetPieceAbbr(piece Piece) string {
	pieceNames := map[Piece]string{
		PAWN:   "P",
		ROOK:   "R",
		KNIGHT: "K",
		BISHOP: "B",
		QUEEN:  "Q",
		KING:   "G",
	}
	return pieceNames[piece]
}
