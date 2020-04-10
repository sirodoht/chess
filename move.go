package main

import (
	"errors"
	"strconv"
	"strings"
)

// Move a command to change the board state
// e.g. "d7 -> d6", so whatever is in d7 put it in d6
type Move struct {
	BeforeLetter rune
	BeforeNumber int
	AfterLetter  rune
	AfterNumber  int
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
func NewMove(command string) (Move, error) {
	// check validity
	if !isMoveSyntaxValid(command) {
		return Move{}, errors.New("invalid move")
	}

	// parse command
	words := strings.Fields(command)
	before := words[0]
	after := words[1]

	m := Move{}
	m.BeforeLetter = []rune(before)[0]
	m.AfterLetter = []rune(after)[0]

	beforeNumber, err := strconv.Atoi(string([]rune(before)[1]))
	if err != nil {
		panic(err)
	}
	m.BeforeNumber = beforeNumber

	afterNumber, err := strconv.Atoi(string([]rune(after)[1]))
	if err != nil {
		panic(err)
	}
	m.AfterNumber = afterNumber

	return m, nil
}

func (m Move) getIndexes(part Part) (int, int) {
	// row
	row := m.AfterNumber - 1
	if part == BEFORE {
		row = m.BeforeNumber - 1
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
	col := columnLetters[m.BeforeLetter]
	if part == AFTER {
		col = columnLetters[m.AfterLetter]
	}

	return row, col
}

// AsString returns the before or after part of the command as string
func (m Move) AsString(part Part) string {
	if part == BEFORE {
		return string(m.BeforeLetter) + strconv.Itoa(m.BeforeNumber)
	}
	return string(m.AfterLetter) + strconv.Itoa(m.AfterNumber)
}

func isMoveSyntaxValid(command string) bool {
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

// getIndexesFromNotation, given a notation like "d7""
// returns the board array's indexes, e.g. (4, 7)
func getIndexesFromNotation(notation string) (int, int) {
	// parse notation
	letter := string([]rune(notation)[0])
	digit := string([]rune(notation)[1])
	digitAsInt, err := strconv.Atoi(digit)
	if err != nil {
		panic(err)
	}

	// row
	row := digitAsInt - 1

	// col
	columnLetters := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
	col := columnLetters[letter]

	return row, col
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
