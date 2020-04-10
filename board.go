package main

import (
	"fmt"
	"os"
	"strings"

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

// Move changes the location of one piece, to a new chess board cell
func (b *Board) Move(command string) {
	// parse command
	words := strings.Fields(command)
	before := words[0]
	after := words[1]

	// cache piece from current location
	oldRow, oldCol := getIndexesFromNotation(before)
	piece := b[oldRow][oldCol]
	pieceName := getPieceName(piece)
	fmt.Printf("\nMOVE: %s commandd to %s\n", pieceName, after)

	// command piece
	newRow, newCol := getIndexesFromNotation(after)
	b[newRow][newCol] = piece
	b[oldRow][oldCol] = "   "
}
