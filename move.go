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
	if !isCommandValid(command) {
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

func (m Move) getIndexes(part Part) (int, int) {
	// row
	row := m.afterNumber - 1
	if part == BEFORE {
		row = m.beforeNumber - 1
	}

	// col
	columnLetters := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
	}
	col := columnLetters[m.beforeLetter]
	if part == AFTER {
		col = columnLetters[m.afterLetter]
	}

	return row, col
}

// AsString returns the before or after part of the command as string
func (m Move) AsString(part Part) string {
	if part == BEFORE {
		return string(m.beforeLetter) + strconv.Itoa(m.beforeNumber)
	}
	return string(m.afterLetter) + strconv.Itoa(m.afterNumber)
}

func isCommandValid(command string) bool {
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
	if !isLetterValid(beforeLetter) {
		return false
	}

	// check after word has valid letter
	afterLetter := []rune(after)[0]
	if !isLetterValid(afterLetter) {
		return false
	}

	// check before word has valid number
	beforeNumber := []rune(before)[1]
	if !isNumberValid(beforeNumber) {
		return false
	}

	// check after word has valid number
	afterNumber := []rune(after)[1]
	if !isNumberValid(afterNumber) {
		return false
	}

	return true
}

func isLetterValid(letter rune) bool {
	validLetters := [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	valid := false
	for _, r := range validLetters {
		if r == letter {
			valid = true
		}
	}

	return valid
}

func isNumberValid(number rune) bool {
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
