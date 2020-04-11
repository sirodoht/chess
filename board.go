package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Board is our chess board state
type Board [8][8]string

// Init adds all pieces in their initial chess positions
func (b *Board) Init() {
	*b = Board{
		{"● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}
}

// Square defines the content of any square on the board
type Square struct {
	team    Team
	isEmpty bool
	piece   Piece
}

// Location is a square on the board.
type Location struct {
	row int
	col int
}

// Execute applies a move to the board
// Essentially, it is the move of a piece on the board.
func (b *Board) Execute(m Move) (string, string) {
	oldRow, oldCol := m.getIndexes(BEFORE)
	content := b[oldRow][oldCol]
	pieceRune := []rune(content)[2]
	piece := GetPiece(pieceRune)
	sq := b.getSquare(m, BEFORE)
	beforeDescription := GetTeamName(m.team) + " " + GetPieceName(sq.piece)

	// command piece
	newRow, newCol := m.getIndexes(AFTER)
	b[newRow][newCol] = GetTeamSymbol(m.team) + " " + GetPieceAbbr(piece)
	b[oldRow][oldCol] = "   "

	return beforeDescription, m.AsString(AFTER)
}

// Render prints the board in stdout
func (b *Board) Render() {
	// init and print board
	table := tablewriter.NewWriter(os.Stdout)

	// add letter headers as tablewriter lib headers
	header := []string{" ", "a", "b", "c", "d", "e", "f", "g", "h"}
	table.SetAutoFormatHeaders(false)
	table.SetHeader(header)

	// transform to [][]string because that's what is required by tablewriter lib
	data := [][]string{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			if len(data) >= i {
				data = append(data, []string{})
			}

			dataCell := ""

			// add row numbers on the left
			if j == 0 {
				dataCell = strconv.Itoa(i + 1)
			} else {
				dataCell = string(b[i][j-1])
			}

			data[i] = append(data[i], dataCell)
		}
	}
	table.AppendBulk(data)
	table.Render()
}

// getSquare returns the part piece that is to be moved, either BEFORE or AFTER
func (b Board) getSquare(m Move, part Part) Square {
	row, col := m.getIndexes(part)
	content := b[row][col]

	color := []rune(content)[0]
	team := WHITE
	if color == '●' {
		team = BLACK
	}

	pieceRune := []rune(content)[1]
	isEmpty := false
	if pieceRune == ' ' {
		isEmpty = true
	}
	piece := PAWN
	pieceSymbols := map[rune]Piece{
		'P': PAWN,
		'R': ROOK,
		'K': KNIGHT,
		'B': BISHOP,
		'Q': QUEEN,
		'G': KING,
	}
	piece = pieceSymbols[pieceRune]

	square := Square{
		team:    team,
		piece:   piece,
		isEmpty: isEmpty,
	}

	return square
}

// getPossibleMoves get a location in the board and a piece,
// and returns all possible location as if the board is empty
func getPossibleMoves(origin Location, piece Piece) []Location {
	possibleMoves := []Location{}
	if piece == ROOK {
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
	}

	return possibleMoves
}

func (b Board) isMoveValid(m Move) bool {
	beforeSquare := b.getSquare(m, BEFORE)
	if beforeSquare.isEmpty {
		return false
	}

	if beforeSquare.piece == ROOK {
		targetSquare := b.getSquare(m, AFTER)
		if beforeSquare.team == targetSquare.team {
			return false
		}
	}
	return true
}
