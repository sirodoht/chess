package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Board is our chess board state
type Board [8][9]string

// Init adds all pieces in their initial chess positions
func (b *Board) Init() {
	*b = Board{
		{"1", "● R", "● K", "● B", "● Q", "● G", "● B", "● K", "● R"},
		{"2", "● P", "● P", "● P", "● P", "● P", "● P", "● P", "● P"},
		{"3", "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"4", "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"5", "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"6", "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		{"7", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P", "○ P"},
		{"8", "○ R", "○ K", "○ B", "○ Q", "○ G", "○ B", "○ K", "○ R"},
	}
}

// Execute applies a move to the board
// Essentially, it is the move of a piece on the board.
func (b *Board) Execute(m Move) (string, string) {
	oldRow, oldCol := m.getIndexes("before")
	piece := b[oldRow][oldCol]
	pieceName := getPieceName(piece)

	// command piece
	newRow, newCol := m.getIndexes("after")
	b[newRow][newCol] = piece
	b[oldRow][oldCol] = "   "

	return pieceName, m.AsString("after")
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
			dataCell := string(b[i][j])
			data[i] = append(data[i], dataCell)
		}
	}
	table.AppendBulk(data)
	table.Render()
}

// getPieceName returns the name of the given piece in notation
// e.g. pieceNotation could be "○ P" and the return string would be "white Pawn"
func getPieceName(pieceNotation string) string {
	// parse piece notation
	circle := []rune(pieceNotation)[0]
	piece := []rune(pieceNotation)[2]

	colorNames := map[rune]string{
		'○': "white ○",
		'●': "black ●",
	}

	pieceNames := map[rune]string{
		'P': "Pawn",
		'R': "Rook",
		'K': "Knight",
		'B': "Bishop",
		'Q': "Queen",
		'G': "King",
	}

	return colorNames[circle] + " " + pieceNames[piece]
}
