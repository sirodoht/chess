package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Move a command to change the board state
// e.g. "d7 -> d6", so whatever is in d7 put it in d6
type Move struct {
	team         Team
	strategy     Strategy
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
	// CHECKMATE is when player checkmates enemy
	CHECKMATE
	// STALEMATE is when non-checked player has no legal move to make
	STALEMATE
)

// NewMove validates and returns a new Move struct out of a command string
// Also returns whether Move was valid, a possible message for user, and whether is
// final move.
func NewMove(b Board, team Team, command string) (Move, bool, []string, bool) {
	if len(command) == 4 {
		command = string(command[0]) + string(command[1]) + " " + string(command[2]) + string(command[3])
	}

	if !IsCommandValid(command) {
		isValid := false
		isEndgame := false
		messages := []string{"MOVE: invalid; example: 'd7 d6'"}
		return Move{}, isValid, messages, isEndgame
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

	m.strategy = GetStrategy(m, b)

	// check move validity
	validityMessage := m.IsValid(b, team)
	if len(validityMessage) > 0 {
		isValid := false
		isEndgame := false
		messages := []string{"MOVE: " + validityMessage}
		return Move{}, isValid, messages, isEndgame
	}

	// build success message
	destinationLocation := m.AsNotation(AFTER)
	originSquare := b.GetSquare(m, BEFORE)
	originPieceName := GetPieceName(originSquare.piece, VERBOSE)
	originTeamName := GetTeamName(m.team, VERBOSE)
	destinationSquare := b.GetSquare(m, AFTER)
	capturedPieceName := GetPieceName(destinationSquare.piece, VERBOSE)
	destinationTeamName := GetTeamName(m.GetEnemy(), SYMBOL)
	msg := fmt.Sprintf("MOVE: %s %s moved to %s", originTeamName, originPieceName, destinationLocation)
	if m.strategy == CAPTURE {
		msg = fmt.Sprintf("CAPTURE: %s %s captured %s %s at %s", originTeamName, originPieceName, destinationTeamName, capturedPieceName, destinationLocation)
	}

	// check if move causes enemy to be in check
	inCheck := IsInCheck(b, m, m.GetEnemy())

	// check if move causes enemy to lose
	checkmated := IsCheckmated(b, m, m.GetEnemy())

	// handle check and checkmate messages
	messages := []string{msg}
	if checkmated {
		checkmateMessage := fmt.Sprintf("CHECKMATE: %s wins!", GetTeamName(m.team, LOWER))
		messages = append(messages, checkmateMessage)
	} else if inCheck {
		checkMessage := fmt.Sprintf("CHECK: %s is in check", destinationTeamName)
		messages = append(messages, checkMessage)
	}

	return m, true, messages, checkmated
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

// GetEnemy return the player color not playing current move
func (m Move) GetEnemy() Team {
	if m.team == WHITE {
		return BLACK
	}
	return WHITE
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
func GetStrategy(m Move, b Board) Strategy {
	beforeSquare := b.GetSquare(m, BEFORE)
	afterSquare := b.GetSquare(m, AFTER)
	if beforeSquare.team != afterSquare.team && !afterSquare.isEmpty {
		return CAPTURE
	}
	return NORMAL
}

// IsValid checks whether the move is valid, given board and whose turn it is
func (m Move) IsValid(b Board, turn Team) string {
	// handle same origin and destination
	if m.GetLocation(BEFORE).row == m.GetLocation(AFTER).row && m.GetLocation(BEFORE).col == m.GetLocation(AFTER).col {
		return "invalid; origin and destination are the same"
	}

	// handle empty square on origin
	beforeSquare := b.GetSquare(m, BEFORE)
	if beforeSquare.isEmpty {
		return "invalid; empty origin"
	}

	// handle when player plays enemy's pieces
	if beforeSquare.team != turn {
		return "invalid; wrong turn"
	}

	// handle when player's destination is same color
	afterSquare := b.GetSquare(m, AFTER)
	if afterSquare.team == turn {
		return "invalid; destination is same color"
	}

	originPiece := beforeSquare.piece
	validity := false
	if originPiece == ROOK {
		validity = m.IsRookMoveValid(b)
	} else if originPiece == KNIGHT {
		validity = m.IsKnightMoveValid(b)
	} else if originPiece == BISHOP {
		validity = m.IsBishopMoveValid(b)
	} else if originPiece == QUEEN {
		validity = m.IsQueenMoveValid(b)
	} else if originPiece == KING {
		validity = m.IsKingMoveValid(b)
	} else if originPiece == PAWN {
		validity = m.IsPawnMoveValid(b)
	}

	if validity == false {
		return "invalid move"
	}

	// check if player playing now is in check
	if IsInCheck(b, m, m.team) {
		return "invalid as checked"
	}

	return ""
}

// GetNotationFromLocation returns string of notation, given Location
func GetNotationFromLocation(location Location) string {
	notationRow := location.row + 1

	colNotations := map[int]rune{
		0: 'a',
		1: 'b',
		2: 'c',
		3: 'd',
		4: 'e',
		5: 'f',
		6: 'g',
		7: 'h',
	}
	notationCol := colNotations[location.col]

	return string(notationCol) + strconv.Itoa(notationRow)
}

// IsInCheck returns true if possiblyCheckedTeam is in check, after given move has been executed
// To find the answer, it scans all board squares, creates moves with each
// enemy piece as origin and current team King as destination, and then checks
// if the move is valid. If so, then that means it's a capture move, which means
// current team's King is in check position.
func IsInCheck(b Board, m Move, possiblyCheckedTeam Team) bool {
	var newBoard Board
	newBoard.LoadData(b)
	newBoard.Execute(m)

	// find attacker team
	attackerTeam := WHITE
	if possiblyCheckedTeam == WHITE {
		attackerTeam = BLACK
	}
	possiblyCheckedKingLocation := newBoard.FindKing(possiblyCheckedTeam)
	possiblyCheckedKingLocationAsNotation := GetNotationFromLocation(possiblyCheckedKingLocation)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			attackerOriginSquare := newBoard.ParseSquare(i, j)

			// omit empty origin squares
			if attackerOriginSquare.isEmpty {
				continue
			}

			// omit moves with origin the possibly checked team
			if attackerOriginSquare.team != attackerTeam {
				continue
			}

			currentLocation := Location{
				row: i,
				col: j,
			}
			if currentLocation.row == possiblyCheckedKingLocation.row && currentLocation.col == possiblyCheckedKingLocation.col {
				continue
			}
			currentLocationAsNotation := GetNotationFromLocation(currentLocation)
			command := currentLocationAsNotation + " " + possiblyCheckedKingLocationAsNotation

			// parse command
			words := strings.Fields(command)
			before := words[0]
			after := words[1]

			// build move to test if it is a check move
			testCheckMove := Move{}
			testCheckMove.team = attackerTeam
			testCheckMove.beforeLetter = []rune(before)[0]
			testCheckMove.afterLetter = []rune(after)[0]
			beforeNumber, err := strconv.Atoi(string([]rune(before)[1]))
			if err != nil {
				panic(err)
			}
			testCheckMove.beforeNumber = beforeNumber
			afterNumber, err := strconv.Atoi(string([]rune(after)[1]))
			if err != nil {
				panic(err)
			}
			testCheckMove.afterNumber = afterNumber
			testCheckMove.strategy = GetStrategy(m, newBoard)

			// validate move
			piece := attackerOriginSquare.piece
			validity := false
			if piece == ROOK {
				validity = testCheckMove.IsRookMoveValid(newBoard)
			} else if piece == KNIGHT {
				validity = testCheckMove.IsKnightMoveValid(newBoard)
			} else if piece == BISHOP {
				validity = testCheckMove.IsBishopMoveValid(newBoard)
			} else if piece == QUEEN {
				validity = testCheckMove.IsQueenMoveValid(newBoard)
			} else if piece == KING {
				validity = testCheckMove.IsKingMoveValid(newBoard)
			} else if piece == PAWN {
				validity = testCheckMove.IsPawnMoveValid(newBoard)
			}

			// if move is valid, then it means King is in check position
			// which means that we should return true
			// if not, then we should continue searching
			if validity {
				return true
			}

		}
	}

	return false
}

// IsCheckmated returns true if given team loses
// It works by finding given team's King and checking if it has any valid moves
func IsCheckmated(b Board, m Move, possiblyCheckmatedTeam Team) bool {
	// create new board and apply current move
	var newBoard Board
	newBoard.LoadData(b)
	newBoard.Execute(m)

	// find all possible destination locations for possibly checkmated King
	possiblyCheckmatedKingLocation := newBoard.FindKing(possiblyCheckmatedTeam)
	kingDestinations := []Location{
		// clockwise
		{row: possiblyCheckmatedKingLocation.row - 1, col: possiblyCheckmatedKingLocation.col},
		{row: possiblyCheckmatedKingLocation.row - 1, col: possiblyCheckmatedKingLocation.col + 1},
		{row: possiblyCheckmatedKingLocation.row, col: possiblyCheckmatedKingLocation.col + 1},
		{row: possiblyCheckmatedKingLocation.row + 1, col: possiblyCheckmatedKingLocation.col + 1},
		{row: possiblyCheckmatedKingLocation.row + 1, col: possiblyCheckmatedKingLocation.col},
		{row: possiblyCheckmatedKingLocation.row + 1, col: possiblyCheckmatedKingLocation.col - 1},
		{row: possiblyCheckmatedKingLocation.row, col: possiblyCheckmatedKingLocation.col - 1},
		{row: possiblyCheckmatedKingLocation.row - 1, col: possiblyCheckmatedKingLocation.col - 1},
	}

	// check if move with origin the King and destination each possible King destination
	// is valid, if no valid are found, then King has been checkmated
	foundValid := false
	for _, currentLocation := range kingDestinations {
		if !IsLocationValid(currentLocation.row, currentLocation.col) {
			continue
		}

		possiblyCheckmatedKingLocationAsNotation := GetNotationFromLocation(possiblyCheckmatedKingLocation)
		currentLocationAsNotation := GetNotationFromLocation(currentLocation)
		command := possiblyCheckmatedKingLocationAsNotation + " " + currentLocationAsNotation

		// parse command
		words := strings.Fields(command)
		before := words[0]
		after := words[1]

		// build move to test if it is a check move
		testKingMove := Move{}
		testKingMove.team = possiblyCheckmatedTeam
		testKingMove.beforeLetter = []rune(before)[0]
		testKingMove.afterLetter = []rune(after)[0]
		beforeNumber, err := strconv.Atoi(string([]rune(before)[1]))
		if err != nil {
			panic(err)
		}
		testKingMove.beforeNumber = beforeNumber
		afterNumber, err := strconv.Atoi(string([]rune(after)[1]))
		if err != nil {
			panic(err)
		}
		testKingMove.afterNumber = afterNumber
		testKingMove.strategy = GetStrategy(m, newBoard)

		// validate move
		isMoveValid := testKingMove.IsKingMoveValid(newBoard)
		if isMoveValid {
			// validate move for check status too
			// if in check, then move not valid
			isMoveValid = !IsInCheck(newBoard, testKingMove, testKingMove.team)
		}

		// if valid move found, then King not checkmated
		if isMoveValid {
			foundValid = true
			break
		}
	}

	// return not foundValid because
	// if valid Move is found, means King is not checkmated (thus return false)
	// if no valid Move is found, means King is checkmated (thus return true)
	return !foundValid
}

// IsRookMoveValid returns whether given move, with Rook as origin piece, is valid
func (m Move) IsRookMoveValid(b Board) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// top
	newRow := originLocation.row - 1
	for IsLocationValid(newRow, originLocation.col) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear, break
				break
			}
			if originLocation.col == destinationLocation.col && newRow == destinationLocation.row {
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear, either found or break
				if originLocation.col == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}
		newRow--
	}

	// bottom
	newRow = originLocation.row + 1
	for IsLocationValid(newRow, originLocation.col) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear, break
				break
			}
			if originLocation.col == destinationLocation.col && newRow == destinationLocation.row {
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, originLocation.col).isEmpty {
				// path is not clear, either found or break
				if originLocation.col == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}
		newRow++
	}

	// left
	newCol := originLocation.col - 1
	for IsLocationValid(originLocation.row, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear, break
				break
			}
			if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear, either found or break
				if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
					return true
				}
				break
			}
		}
		newCol--
	}

	// right
	newCol = originLocation.col + 1
	for IsLocationValid(originLocation.row, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear, break
				break
			}
			if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(originLocation.row, newCol).isEmpty {
				// path is not clear, either found or break
				if originLocation.row == destinationLocation.row && newCol == destinationLocation.col {
					return true
				}
				break
			}
		}
		newCol++
	}

	return false
}

// IsKnightMoveValid returns whether given move, with Knight as origin piece, is valid
func (m Move) IsKnightMoveValid(b Board) bool {
	// searching for Knight moves in the fashion of
	// two hops forward, then one left, or one right
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// handle top hand
	newRow := originLocation.row - 2
	newCol := originLocation.col - 1 // left side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newCol = originLocation.col + 1 // right side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle right hand
	newCol = originLocation.col + 2
	newRow = originLocation.row - 1 // top side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newRow = originLocation.row + 1 // bottom side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle bottom hand
	newRow = originLocation.row + 2
	newCol = originLocation.col - 1 // left side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newCol = originLocation.col + 1 // right side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// handle left hand
	newCol = originLocation.col - 2
	newRow = originLocation.row - 1 // top side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}
	newRow = originLocation.row + 1 // bottom side
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	return false
}

// IsBishopMoveValid returns whether given move, with Bishop as origin piece, is valid
func (m Move) IsBishopMoveValid(b Board) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// go top-right
	newRow := originLocation.row - 1
	newCol := originLocation.col + 1
	for IsLocationValid(newRow, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
			if newRow == destinationLocation.row && newCol == destinationLocation.col {
				// reached destination
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear, either found or break
				if newCol == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}

		newRow = newRow - 1
		newCol = newCol + 1
	}

	// go bottom-right
	newRow = originLocation.row + 1
	newCol = originLocation.col + 1
	for IsLocationValid(newRow, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
			if newRow == destinationLocation.row && newCol == destinationLocation.col {
				// reached destination
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear, either found or break
				if newCol == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}
		newRow = newRow + 1
		newCol = newCol + 1
	}

	// go bottom-left
	newRow = originLocation.row + 1
	newCol = originLocation.col - 1
	for IsLocationValid(newRow, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
			if newRow == destinationLocation.row && newCol == destinationLocation.col {
				// reached destination
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear, either found or break
				if newCol == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}
		newRow = newRow + 1
		newCol = newCol - 1
	}

	// go top-left
	newRow = originLocation.row - 1
	newCol = originLocation.col - 1
	for IsLocationValid(newRow, newCol) {
		if m.strategy == NORMAL {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear
				break
			}
			if newRow == destinationLocation.row && newCol == destinationLocation.col {
				// reached destination
				return true
			}
		} else if m.strategy == CAPTURE {
			if !b.ParseSquare(newRow, newCol).isEmpty {
				// path is not clear, either found or break
				if newCol == destinationLocation.col && newRow == destinationLocation.row {
					return true
				}
				break
			}
		}
		newRow = newRow - 1
		newCol = newCol - 1
	}

	return false
}

// IsQueenMoveValid returns whether given move, with Queen as origin piece, is valid
func (m Move) IsQueenMoveValid(b Board) bool {
	rookMovesValidity := m.IsRookMoveValid(b)
	bishopMovesValidity := m.IsBishopMoveValid(b)

	return rookMovesValidity || bishopMovesValidity
}

// IsKingMoveValid returns whether given move, with King as origin piece, is valid
func (m Move) IsKingMoveValid(b Board) bool {
	originLocation := m.GetLocation(BEFORE)
	destinationLocation := m.GetLocation(AFTER)

	// vertical top
	newRow := originLocation.row - 1
	newCol := originLocation.col
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal top-right
	newRow = originLocation.row - 1
	newCol = originLocation.col + 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// horizontal right
	newRow = originLocation.row
	newCol = originLocation.col + 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal bottom-right
	newRow = originLocation.row + 1
	newCol = originLocation.col + 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// vertical bottom
	newRow = originLocation.row + 1
	newCol = originLocation.col
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal bottom-left
	newRow = originLocation.row + 1
	newCol = originLocation.col - 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// horizontal left
	newRow = originLocation.row
	newCol = originLocation.col - 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	// diagonal top-left
	newRow = originLocation.row - 1
	newCol = originLocation.col - 1
	if m.strategy == NORMAL || m.strategy == CAPTURE {
		if newRow == destinationLocation.row && newCol == destinationLocation.col {
			return true
		}
	}

	return false
}

// IsPawnMoveValid returns whether given move, with Pawn as origin piece, is valid
func (m Move) IsPawnMoveValid(b Board) bool {
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
		if m.strategy == NORMAL {
			if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
				return true
			}
		} else if m.strategy == CAPTURE {
			newColLeft := originLocation.col - 1
			if newRow == destinationLocation.row && newColLeft == destinationLocation.col {
				return true
			}
			newColRight := originLocation.col + 1
			if newRow == destinationLocation.row && newColRight == destinationLocation.col {
				return true
			}
		}
		if firstMove {
			newRow--
			if m.strategy == NORMAL {
				if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
					return true
				}
			}
		}
	}

	// if black / up side
	if m.team == BLACK {
		newRow := originLocation.row + 1
		if m.strategy == NORMAL {
			if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
				return true
			}
		} else if m.strategy == CAPTURE {
			newColLeft := originLocation.col + 1
			if newRow == destinationLocation.row && newColLeft == destinationLocation.col {
				return true
			}
			newColRight := originLocation.col - 1
			if newRow == destinationLocation.row && newColRight == destinationLocation.col {
				return true
			}
		}
		if firstMove {
			newRow++
			if m.strategy == NORMAL {
				if newRow == destinationLocation.row && originLocation.col == destinationLocation.col {
					return true
				}
			}
		}
	}

	return false
}
