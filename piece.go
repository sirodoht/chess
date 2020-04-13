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
	} else {
		upperPieceNames := map[Piece]string{
			PAWN:   "PAWN",
			ROOK:   "ROOK",
			KNIGHT: "KNIGHT",
			BISHOP: "BISHOP",
			QUEEN:  "QUEEN",
			KING:   "KING",
		}
		return upperPieceNames[piece]
	}
}

// GetPossibleMoves get a location in the board and a piece,
// and returns all possible location as if the board is empty
func GetPossibleMoves(origin Location, piece Piece) []Location {
	possibleMoves := []Location{}

	if piece == ROOK {
		possibleMoves = GetPossibleRookMoves(origin)
	} else if piece == KNIGHT {
		possibleMoves = GetPossibleKnightMoves(origin)
	} else if piece == BISHOP {
		possibleMoves = GetPossibleBishopMoves(origin)
	} else if piece == QUEEN {
		possibleMoves = GetPossibleQueenMoves(origin)
	} else if piece == KING {
		possibleMoves = GetPossibleKingMoves(origin)
	} else if piece == PAWN {
		possibleMoves = GetPossiblePawnMoves(origin)
	}

	return possibleMoves
}

// AddPossibleMove adds a possible move to given possibleMoves slice
// after validating that row and col are inside board.
// Also returns whether the move was inserted.
func AddPossibleMove(row int, col int, possibleMoves []Location) ([]Location, bool) {
	inserted := false
	if row >= 0 && row <= 7 && col >= 0 && col <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: row,
			col: col,
		})
		inserted = true
	}
	return possibleMoves, inserted
}

// GetPossibleRookMoves returns all possible moves for a Rook,
// given origin current location on the board
func GetPossibleRookMoves(origin Location) []Location {
	possibleMoves := []Location{}
	var inserted bool

	// top
	newRow := origin.row - 1
	for {
		possibleMoves, inserted = AddPossibleMove(newRow, origin.col, possibleMoves)
		if !inserted {
			break
		}
		newRow--
	}

	// bottom
	newRow = origin.row + 1
	for {
		possibleMoves, inserted = AddPossibleMove(newRow, origin.col, possibleMoves)
		if !inserted {
			break
		}
		newRow++
	}

	// left
	newCol := origin.col - 1
	for {
		possibleMoves, inserted = AddPossibleMove(origin.row, newCol, possibleMoves)
		if !inserted {
			break
		}
		newCol--
	}

	// right
	newCol = origin.col + 1
	for {
		possibleMoves, inserted = AddPossibleMove(origin.row, newCol, possibleMoves)
		if !inserted {
			break
		}
		newCol++
	}

	return possibleMoves
}

// GetPossibleKnightMoves returns all possible moves for a Knight,
// given origin current location on the board
func GetPossibleKnightMoves(origin Location) []Location {
	// searching for Knight moves in a fashion
	// of two hops forward, then one left, or one right
	possibleMoves := []Location{}

	// handle top hand -> left side
	newRow := origin.row - 2
	newCol := origin.col - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle top hand -> right side
	newCol = origin.col + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle right hand -> top side
	newCol = origin.col + 2
	newRow = origin.row - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle right hand -> bottom side
	newRow = origin.row + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle bottom hand -> left side
	newRow = origin.row + 2
	newCol = origin.col - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle bottom hand -> right side
	newCol = origin.col + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle left hand -> top side
	newCol = origin.col - 2
	newRow = origin.row - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// handle left hand -> botom side
	newRow = origin.row + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	return possibleMoves
}

// GetPossibleBishopMoves returns all possible moves for a Bishop,
// given origin current location on the board
func GetPossibleBishopMoves(origin Location) []Location {
	possibleMoves := []Location{}

	// go top-right
	newRow := origin.row
	newCol := origin.col
	inserted := true
	for inserted {
		newRow = newRow - 1
		newCol = newCol + 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// go bottom-right
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow + 1
		newCol = newCol + 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// go bottom-left
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow + 1
		newCol = newCol - 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// go top-left
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow - 1
		newCol = newCol - 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	return possibleMoves
}

// GetPossibleQueenMoves returns all possible moves for a Queen,
// given origin current location on the board
func GetPossibleQueenMoves(origin Location) []Location {
	possibleMoves := []Location{}
	var inserted bool

	// top
	newRow := origin.row - 1
	for {
		possibleMoves, inserted = AddPossibleMove(newRow, origin.col, possibleMoves)
		if !inserted {
			break
		}
		newRow--
	}

	// bottom
	newRow = origin.row + 1
	for {
		possibleMoves, inserted = AddPossibleMove(newRow, origin.col, possibleMoves)
		if !inserted {
			break
		}
		newRow++
	}

	// left
	newCol := origin.col - 1
	for {
		possibleMoves, inserted = AddPossibleMove(origin.row, newCol, possibleMoves)
		if !inserted {
			break
		}
		newCol--
	}

	// right
	newCol = origin.col + 1
	for {
		possibleMoves, inserted = AddPossibleMove(origin.row, newCol, possibleMoves)
		if !inserted {
			break
		}
		newCol++
	}

	// diagonal top-right
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow - 1
		newCol = newCol + 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// diagonal bottom-right
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow + 1
		newCol = newCol + 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// diagonal bottom-left
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow + 1
		newCol = newCol - 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	// diagonal top-left
	newRow = origin.row
	newCol = origin.col
	inserted = true
	for inserted {
		newRow = newRow - 1
		newCol = newCol - 1
		possibleMoves, inserted = AddPossibleMove(newRow, newCol, possibleMoves)
	}

	return possibleMoves
}

// GetPossibleKingMoves returns all possible moves for a King,
// given origin current location on the board
func GetPossibleKingMoves(origin Location) []Location {
	possibleMoves := []Location{}

	// vertical top
	newRow := origin.row - 1
	newCol := origin.col
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// diagonal top-right
	newRow = origin.row - 1
	newCol = origin.col + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// horizontal right
	newRow = origin.row
	newCol = origin.col + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// diagonal bottom-right
	newRow = origin.row + 1
	newCol = origin.col + 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// vertical bottom
	newRow = origin.row + 1
	newCol = origin.col
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// diagonal bottom-left
	newRow = origin.row + 1
	newCol = origin.col - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// horizontal left
	newRow = origin.row
	newCol = origin.col - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// diagonal top-left
	newRow = origin.row - 1
	newCol = origin.col - 1
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	return possibleMoves
}

// GetPossiblePawnMoves returns all possible moves for a Pawn,
// given origin current location on the board, without knowing if its white or black.
func GetPossiblePawnMoves(origin Location) []Location {
	possibleMoves := []Location{}

	// find out if pawn is on first move
	firstMove := false
	if origin.row == 1 || origin.row == 6 {
		firstMove = true
	}

	// if white / down side
	newRow := origin.row - 1
	possibleMoves, _ = AddPossibleMove(newRow, origin.col, possibleMoves)
	if firstMove {
		newRow--
		possibleMoves, _ = AddPossibleMove(newRow, origin.col, possibleMoves)
	}

	// if white / down and capture opponents piece
	newRow = origin.row - 1
	newCol := origin.col - 1 // left side
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)
	newCol = origin.col + 1 // right side
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	// if black / up side
	newRow = origin.row + 1
	possibleMoves, _ = AddPossibleMove(newRow, origin.col, possibleMoves)
	if firstMove {
		newRow++
		possibleMoves, _ = AddPossibleMove(newRow, origin.col, possibleMoves)
	}

	// if black / up and capture opponents piece
	newRow = origin.row + 1
	newCol = origin.col - 1 // left side
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)
	newCol = origin.col + 1 // right side
	possibleMoves, _ = AddPossibleMove(newRow, newCol, possibleMoves)

	return possibleMoves
}
