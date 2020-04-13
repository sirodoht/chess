package main

import (
	"errors"
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

// NewMove returns a new Move struct out of a command string
func NewMove(team Team, command string) (Move, error) {
	// check command validity
	if !IsCommandValid(command) {
		return Move{}, errors.New("invalid move")
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
