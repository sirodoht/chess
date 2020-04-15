package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Move a command to change the board state
// e.g. "d7 -> d6", so whatever is in d7 put it in d6
type Move struct {
	team         Team
	beforeLetter rune
	beforeNumber int
	afterLetter  rune
	afterNumber  int
}

// Part defines either the before or the after part of a move
// e.g. if move is "d7 d6" the before is "d7" and the after is "d6"
type Part int

const (
	// BEFORE is about the part of the move that defines where the piece comes from
	BEFORE Part = iota

	// AFTER is about the part of the move that defines where the piece goes to
	AFTER
)

// Strategy defines whether a move is a capture one, or a simple one,
// or a castling, etc.
type Strategy int

const (
	// NORMAL is when player moves piece to an empty square
	NORMAL Strategy = iota

	// CAPTURE is when player moves piece to capture enemy piece
	CAPTURE

	// CASTLING is when player executes a castling between King and Rook
	CASTLING

	// ENPASSANT is when player captures enemy pawn in passing
	ENPASSANT

	// PROMOTION is when pawn reaches eighth square and is promoted
	PROMOTION

	// CHECK is when player threatens enemy's King
	CHECK

	// CHECKMATE is when player checkmates enemy
	CHECKMATE

	// STALEMATE is when non-checked player has no legal move to make
	STALEMATE
)

// NewMove validates and returns a new Move struct out of a command string
func NewMove(b Board, team Team, command string) (Move, error) {
	if len(command) == 4 {
		command = string(command[0]) + string(command[1]) + " " + string(command[2]) + string(command[3])
	}

	if !IsCommandValid(command) {
		return Move{}, errors.New("invalid command")
	}

	// parse command
	words := strings.Fields(command)
	before := words[0]
	after := words[1]

	m := Move{}
	m.team = team
	m.beforeLetter = []rune(before)[0]
	m.afterLetter = []rune(after)[0]

	beforeNumber, err := strconv.Atoi(string([]rune(before)[1]))
	if err != nil {
		panic(err)
	}
	m.beforeNumber = beforeNumber

	afterNumber, err := strconv.Atoi(string([]rune(after)[1]))
	if err != nil {
		panic(err)
	}
	m.afterNumber = afterNumber

	// check move validity
	if !m.IsValid(b, team) {
		return Move{}, errors.New("invalid move")
	}

	return m, nil
}

// GetLocation returns the Location struct of either BEFORE or AFTER parts
func (m Move) GetLocation(part Part) Location {
	// row
	row := m.afterNumber - 1
	if part == BEFORE {
		row = m.beforeNumber - 1
	}

	// col
	columnLetters := map[rune]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}
	col := columnLetters[m.beforeLetter]
	if part == AFTER {
		col = columnLetters[m.afterLetter]
	}

	return Location{
		row: row,
		col: col,
	}
}

// AsNotation returns the before or after part of the command as chess notation
// e.g. d7
func (m Move) AsNotation(part Part) string {
	if part == BEFORE {
		return string(m.beforeLetter) + strconv.Itoa(m.beforeNumber)
	}
	return string(m.afterLetter) + strconv.Itoa(m.afterNumber)
}

// IsCommandValid returns whether a command is valid
// A command is a chess move notation
// e.g. "d7 d6", which means piece that is at d7, should go to d6
func IsCommandValid(command string) bool {
	words := strings.Fields(command)

	// check that there are two words
	if len(words) != 2 {
		return false
	}

	before := words[0]
	after := words[1]

	// check both words have 2 letters
	if len(before) != 2 {
		return false
	}
	if len(after) != 2 {
		return false
	}

	// check before word has valid letter
	beforeLetter := []rune(before)[0]
	if !IsLetterValid(beforeLetter) {
		return false
	}

	// check after word has valid letter
	afterLetter := []rune(after)[0]
	if !IsLetterValid(afterLetter) {
		return false
	}

	// check before word has valid number
	beforeNumber := []rune(before)[1]
	if !IsNumberValid(beforeNumber) {
		return false
	}

	// check after word has valid number
	afterNumber := []rune(after)[1]
	if !IsNumberValid(afterNumber) {
		return false
	}

	return true
}

// IsLetterValid returns whether inputted letter in the command is valid
func IsLetterValid(letter rune) bool {
	validLetters := [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	valid := false
	for _, r := range validLetters {
		if r == letter {
			valid = true
		}
	}

	return valid
}

// IsNumberValid returns whether inputted number in the command is valid
func IsNumberValid(number rune) bool {
	numberStr := string(number)
	validNumbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	beforeNumberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		panic(err)
	}
	valid := false
	for _, n := range validNumbers {
		if n == beforeNumberInt {
			valid = true
		}
	}
	return valid
}

// GetStrategy identifies what strategy player goes for
func (m Move) GetStrategy(b Board) Strategy {
	afterSquare := b.GetSquare(m, AFTER)
	if afterSquare.isEmpty {
		return NORMAL
	}
	return CAPTURE
}

// IsValid checks whether the move is valid, given board and whose turn it is
func (m Move) IsValid(b Board, turn Team) bool {
	// handle empty square on origin
	beforeSquare := b.GetSquare(m, BEFORE)
	if beforeSquare.isEmpty {
		fmt.Println("empty origin")
		return false
	}

	// handle when player plays enemy's pieces
	if beforeSquare.team != turn {
		fmt.Printf("wrong turn")
		return false
	}

	// handle when player's destination is same color
	afterSquare := b.GetSquare(m, AFTER)
	if afterSquare.team == turn {
		fmt.Printf("destination is same color")
		return false
	}

	originPiece := beforeSquare.piece
	strategy := m.GetStrategy(b)
	if originPiece == ROOK {
		return m.IsRookMoveValid(b, strategy)
	} else if originPiece == KNIGHT {
		return m.IsKnightMoveValid(b, strategy)
	} else if originPiece == BISHOP {
		return m.IsBishopMoveValid(b, strategy)
	} else if originPiece == QUEEN {
		return m.IsQueenMoveValid(b, strategy)
	} else if originPiece == KING {
		return m.IsKingMoveValid(b, strategy)
	} else if originPiece == PAWN {
		return m.IsPawnMoveValid(b, strategy)
	}

	return false
}

// IsRookMoveValid returns whether given move, with Rook as origin piece, is valid
func (m Move) IsRookMoveValid(b Board, strategy Strategy) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// top
	newRow := originLocation.row - 1
	for IsLocationValid(newRow, originLocation.col) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear
				break
			}
		}
		if destinationLocation.col == originLocation.col && newRow == destinationLocation.row {
			if strategy == NORMAL {
				return true
			}
		}
		newRow--
	}

	// bottom
	newRow = originLocation.row + 1
	for IsLocationValid(newRow, originLocation.col) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear
				break
			}
		}
		if destinationLocation.col == originLocation.col && newRow == destinationLocation.row {
			if strategy == NORMAL {
				return true
			}
		}
		newRow++
	}

	// left
	newCol := originLocation.col - 1
	for IsLocationValid(originLocation.row, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
			if strategy == NORMAL {
				return true
			}
		}
		newCol--
	}

	// right
	newCol = originLocation.col + 1
	for IsLocationValid(originLocation.row, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
			if strategy == NORMAL {
				return true
			}
		}
		newCol++
	}

	return false
}

// IsKnightMoveValid returns whether given move, with Knight as origin piece, is valid
func (m Move) IsKnightMoveValid(b Board, strategy Strategy) bool {
	// searching for Knight moves in the fashion of
	// two hops forward, then one left, or one right
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// handle top hand
	newRow := originLocation.row - 2
	newCol := originLocation.col - 1 // left side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newCol = originLocation.col + 1 // right side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle right hand
	newCol = originLocation.col + 2
	newRow = originLocation.row - 1 // top side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newRow = originLocation.row + 1 // bottom side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle bottom hand
	newRow = originLocation.row + 2
	newCol = originLocation.col - 1 // left side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newCol = originLocation.col + 1 // right side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle left hand
	newCol = originLocation.col - 2
	newRow = originLocation.row - 1 // top side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newRow = originLocation.row + 1 // bottom side
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	return false
}

// IsBishopMoveValid returns whether given move, with Bishop as origin piece, is valid
func (m Move) IsBishopMoveValid(b Board, strategy Strategy) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// go top-right
	newRow := originLocation.row - 1
	newCol := originLocation.col + 1
	for IsLocationValid(newRow, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			// reached destination
			if strategy == NORMAL {
				return true
			}
		}
		newRow = newRow - 1
		newCol = newCol + 1
	}

	// go bottom-right
	newRow = originLocation.row + 1
	newCol = originLocation.col + 1
	for IsLocationValid(newRow, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			// reached destination
			if strategy == NORMAL {
				return true
			}
		}
		newRow = newRow + 1
		newCol = newCol + 1
	}

	// go bottom-left
	newRow = originLocation.row + 1
	newCol = originLocation.col - 1
	for IsLocationValid(newRow, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			// reached destination
			if strategy == NORMAL {
				return true
			}
		}
		newRow = newRow + 1
		newCol = newCol - 1
	}

	// go top-left
	newRow = originLocation.row - 1
	newCol = originLocation.col - 1
	for IsLocationValid(newRow, newCol) {
		if strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
		}
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			// reached destination
			if strategy == NORMAL {
				return true
			}
		}
		newRow = newRow - 1
		newCol = newCol - 1
	}

	return false
}

// IsQueenMoveValid returns whether given move, with Queen as origin piece, is valid
func (m Move) IsQueenMoveValid(b Board, strategy Strategy) bool {
	rookMovesValidity := m.IsRookMoveValid(b, strategy)
	bishopMovesValidity := m.IsBishopMoveValid(b, strategy)

	return rookMovesValidity || bishopMovesValidity
}

// IsKingMoveValid returns whether given move, with King as origin piece, is valid
func (m Move) IsKingMoveValid(b Board, strategy Strategy) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// vertical top
	newRow := originLocation.row - 1
	newCol := originLocation.col
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal top-right
	newRow = originLocation.row - 1
	newCol = originLocation.col + 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// horizontal right
	newRow = originLocation.row
	newCol = originLocation.col + 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal bottom-right
	newRow = originLocation.row + 1
	newCol = originLocation.col + 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// vertical bottom
	newRow = originLocation.row + 1
	newCol = originLocation.col
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal bottom-left
	newRow = originLocation.row + 1
	newCol = originLocation.col - 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// horizontal left
	newRow = originLocation.row
	newCol = originLocation.col - 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal top-left
	newRow = originLocation.row - 1
	newCol = originLocation.col - 1
	if strategy == NORMAL {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	return false
}

// IsPawnMoveValid returns whether given move, with Pawn as origin piece, is valid
func (m Move) IsPawnMoveValid(b Board, strategy Strategy) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// find out if pawn is on first move
	firstMove := false
	if (originLocation.row == 1 && m.team == BLACK) || (originLocation.row == 6 && m.team == WHITE) {
		firstMove = true
	}

	// if white / down side
	if m.team == WHITE {
		newRow := originLocation.row - 1
		if strategy == NORMAL {
			if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
				return true
			}
		}
		if firstMove {
			newRow--
			if strategy == NORMAL {
				if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
					return true
				}
			}
		}
	}

	// if black / up side
	if m.team == BLACK {
		newRow := originLocation.row + 1
		if strategy == NORMAL {
			if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
				return true
			}
		}
		if firstMove {
			newRow++
			if strategy == NORMAL {
				if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
					return true
				}
			}
		}
	}

	return false
}
