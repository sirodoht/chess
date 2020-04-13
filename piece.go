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

	// vertical, all rows, same column
	newRow := 0
	for newRow < 8 {
		if newRow == origin.row {
			// omit current location
			newRow++
			continue
		}
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: origin.col,
		})
		newRow++
	}

	// horizontal, all columns, same row
	newCol := 0
	for newCol < 8 {
		if newCol == origin.col {
			// omit current location
			newCol++
			continue
		}
		possibleMoves = append(possibleMoves, Location{
			row: origin.row,
			col: newCol,
		})
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

	// handle top hand
	newRow := origin.row - 2
	if newRow >= 0 {
		newCol := origin.col - 1
		// check for left side
		if newCol >= 0 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}
		newCol = origin.col + 1
		// check for right side
		if newCol <= 7 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}
	}

	// handle right hand
	newCol := origin.col + 2
	if newCol <= 7 {
		newRow := origin.row - 1
		// check for top side
		if newRow >= 0 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}

		newRow = origin.row + 1
		// check for bottom side
		if newRow <= 7 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}
	}

	// handle bottom hand
	newRow = origin.row + 2
	if newRow <= 7 {
		newCol := origin.col - 1
		// check for left side
		if newCol >= 0 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}

		newCol = origin.col + 1
		// check for right side
		if newCol <= 7 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}
	}

	// handle left hand
	newCol = origin.col - 2
	if newCol >= 0 {
		newRow := origin.row - 1
		// check for top side
		if newRow <= 7 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}

		newRow = origin.row + 1
		// check for botom side
		if newRow >= 0 {
			possibleMoves = append(possibleMoves, Location{
				row: newRow,
				col: newCol,
			})
		}
	}

	return possibleMoves
}

// GetPossibleBishopMoves returns all possible moves for a Bishop,
// given origin current location on the board
func GetPossibleBishopMoves(origin Location) []Location {
	possibleMoves := []Location{}

	// go top-right
	newRow := origin.row - 1
	newCol := origin.col + 1
	for newRow >= 0 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow - 1
		newCol = newCol + 1
	}

	// go bottom-right
	newRow = origin.row + 1
	newCol = origin.col + 1
	for newRow <= 7 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow + 1
		newCol = newCol + 1
	}

	// go bottom-left
	newRow = origin.row + 1
	newCol = origin.col - 1
	for newRow <= 7 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow + 1
		newCol = newCol - 1
	}

	// go top-left
	newRow = origin.row - 1
	newCol = origin.col - 1
	for newRow >= 0 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow - 1
		newCol = newCol - 1
	}

	return possibleMoves
}

// GetPossibleQueenMoves returns all possible moves for a Queen,
// given origin current location on the board
func GetPossibleQueenMoves(origin Location) []Location {
	possibleMoves := []Location{}

	// vertical, all rows, same column
	newRow := 0
	for newRow < 8 {
		if newRow == origin.row {
			// omit current location
			newRow++
			continue
		}
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: origin.col,
		})
		newRow++
	}

	// horizontal, all columns, same row
	newCol := 0
	for newCol < 8 {
		if newCol == origin.col {
			// omit current location
			newCol++
			continue
		}
		possibleMoves = append(possibleMoves, Location{
			row: origin.row,
			col: newCol,
		})
		newCol++
	}

	// diagonal top-right
	newRow = origin.row - 1
	newCol = origin.col + 1
	for newRow >= 0 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow - 1
		newCol = newCol + 1
	}

	// diagonal bottom-right
	newRow = origin.row + 1
	newCol = origin.col + 1
	for newRow <= 7 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow + 1
		newCol = newCol + 1
	}

	// diagonal bottom-left
	newRow = origin.row + 1
	newCol = origin.col - 1
	for newRow <= 7 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow + 1
		newCol = newCol - 1
	}

	// diagonal top-left
	newRow = origin.row - 1
	newCol = origin.col - 1
	for newRow >= 0 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
		newRow = newRow - 1
		newCol = newCol - 1
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
	if newRow >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// diagonal top-right
	newRow = origin.row - 1
	newCol = origin.col + 1
	if newRow >= 0 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// horizontal right
	newRow = origin.row
	newCol = origin.col + 1
	if newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// diagonal bottom-right
	newRow = origin.row + 1
	newCol = origin.col + 1
	if newRow <= 7 && newCol <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// vertical bottom
	newRow = origin.row + 1
	newCol = origin.col
	if newRow <= 7 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// diagonal bottom-left
	newRow = origin.row + 1
	newCol = origin.col - 1
	if newRow <= 7 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// horizontal left
	newRow = origin.row
	newCol = origin.col - 1
	if newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

	// diagonal top-left
	newRow = origin.row - 1
	newCol = origin.col - 1
	if newRow >= 0 && newCol >= 0 {
		possibleMoves = append(possibleMoves, Location{
			row: newRow,
			col: newCol,
		})
	}

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
