package main

// Team is the white or black team / player of a chess game
type Team int

const (
	// WHITE is the white team / player on the chess board
	WHITE Team = iota
	// BLACK is the black team / player on the chess board
	BLACK
)
