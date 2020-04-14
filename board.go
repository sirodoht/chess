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
	// change piece position on the board
	oldLocation := m.GetLocation(BEFORE)
	newLocation := m.GetLocation(AFTER)
	square := b.GetSquare(m, BEFORE)
	b[newLocation.row][newLocation.col] = GetTeamName(m.team, SYMBOL) + " " + GetPieceName(square.piece, SYMBOL)
	b[oldLocation.row][oldLocation.col] = "   "

	beforeDescription := GetTeamName(m.team, VERBOSE) + " " + GetPieceName(square.piece, VERBOSE)
	return beforeDescription, m.AsNotation(AFTER)
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

// GetSquare returns the part piece that is to be moved, either BEFORE or AFTER
func (b Board) GetSquare(m Move, part Part) Square {
	location := m.GetLocation(part)
	content := b[location.row][location.col]

	color := []rune(content)[0]
	team := NEITHER
	if color == '●' {
		team = BLACK
	} else if color == '○' {
		team = WHITE
	}

	pieceRune := []rune(content)[2]
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

// ParseSquare returns square based on indexes
func (b Board) ParseSquare(row int, col int) Square {
	content := b[row][col]

	color := []rune(content)[0]
	team := NEITHER
	if color == '●' {
		team = BLACK
	} else if color == '○' {
		team = WHITE
	}

	pieceRune := []rune(content)[2]
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
