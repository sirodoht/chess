package main

// Format is in what format the various structs can be returned
// These structs are Team, Piece
type Format int

const (
	// VERBOSE is e.g. "white ○"
	VERBOSE Format = iota
	// SYMBOL is only the circle symbol, e.g. "○"
	SYMBOL
	// UPPER is as uppercase letters
	UPPER
)
