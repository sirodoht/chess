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

// GetPieceName returns given piece in given format
func GetPieceName(piece Piece, format Format) string {
	if format == SYMBOL {
		symbolPieceNames := map[Piece]string{
			PAWN:   "P",
			ROOK:   "R",
			KNIGHT: "K",
			BISHOP: "B",
			QUEEN:  "Q",
			KING:   "G",
		}
		return symbolPieceNames[piece]
	} else if format == VERBOSE {
		verbosePieceNames := map[Piece]string{
			PAWN:   "Pawn",
			ROOK:   "Rook",
			KNIGHT: "Knight",
			BISHOP: "Bishop",
			QUEEN:  "Queen",
			KING:   "King",
		}
		return verbosePieceNames[piece]
	} else if format == UPPER {
		upperPieceNames := map[Piece]string{
			PAWN:   "PAWN",
			ROOK:   "ROOK",
			KNIGHT: "KNIGHT",
			BISHOP: "BISHOP",
			QUEEN:  "QUEEN",
			KING:   "KING",
		}
		return upperPieceNames[piece]
	} else {
		lowerPieceNames := map[Piece]string{
			PAWN:   "pawn",
			ROOK:   "rook",
			KNIGHT: "knight",
			BISHOP: "bishop",
			QUEEN:  "queen",
			KING:   "king",
		}
		return lowerPieceNames[piece]
	}
}
