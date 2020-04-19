package main

import (
	"errors"
	"fmt"
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

// NewLocation validates and returns a new Location struct
func NewLocation(row int, col int) (Location, error) {
	if row >= 0 && row <= 7 && col >= 0 && col <= 7 {
		return Location{
			row: row,
			col: col,
		}, nil
	}
	return Location{}, errors.New("invalid row/col")
}

// IsLocationValid validates location row and col
func IsLocationValid(row int, col int) bool {
	if row >= 0 && row <= 7 && col >= 0 && col <= 7 {
		return true
	}
	return false
}

// Execute applies a move to the board
// Essentially, it is the move of a piece on the board.
func (b *Board) Execute(m Move) {
	// change piece position on the board
	oldLocation := m.GetLocation(BEFORE)
	newLocation := m.GetLocation(AFTER)
	square := b.GetSquare(m, BEFORE)

	if !IsLocationValid(newLocation.row, newLocation.col) {
		errorMsg := fmt.Sprintf("move destination location (%d:%d) is invalid", newLocation.row, newLocation.col)
		panic(errorMsg)
	}
	b[newLocation.row][newLocation.col] = GetTeamName(m.team, SYMBOL) + " " + GetPieceName(square.piece, SYMBOL)

	if !IsLocationValid(oldLocation.row, oldLocation.col) {
		errorMsg := fmt.Sprintf("move origin location (%d:%d) is invalid", oldLocation.row, oldLocation.col)
		panic(errorMsg)
	}
	b[oldLocation.row][oldLocation.col] = "   "
}

// Render prints the board in stdout
func (b *Board) Render() {
	fmt.Println() // to breathe
	fmt.Println("   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |")
	fmt.Println(" - +-----+-----+-----+-----+-----+-----+-----+-----+")

	for i := 0; i < 8; i++ {
		fmt.Printf(" %d |", i)
		for j := 0; j < 8; j++ {
			cell := b[i][j]
			fmt.Printf(" %s |", cell)
		}
		fmt.Println()
	}

	fmt.Println(" - +-----+-----+-----+-----+-----+-----+-----+-----+")
	fmt.Println() // breathe again
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

// FindKing returns the square of the King on current board
func (b Board) FindKing(team Team) Location {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			square := b.ParseSquare(i, j)
			if square.team == team && square.piece == KING {
				return Location{
					row: i,
					col: j,
				}
			}
		}
	}
	return Location{}
}

// LoadData loads all data into board from another board
func (b *Board) LoadData(originBoard Board) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b[i][j] = originBoard[i][j]
		}
	}
}
