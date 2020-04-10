package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// main game loop
	for {
		// initialize board
		board := Board{}
		board.Init()
		board.Render()

		// read from terminal
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter next move: ")
		move, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		move = strings.TrimSpace(move)

		// check move syntax validity
		if !isMoveSyntaxValid(move) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		// check move validity in current board
		// if !isMoveValid(move) {
		// 	fmt.Println("Invalid move. Please try again.")
		// 	continue
		// }

		// execute move
		board.Move(move)

		// check for exit
		if move == "exit" {
			break
		}
	}
}

func isMoveSyntaxValid(move string) bool {
	words := strings.Fields(move)

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

	validLetters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	// check before word has valid letters
	beforeLetter := []rune(before)[0]
	valid := false
	for _, r := range validLetters {
		if r == beforeLetter {
			valid = true
		}
	}
	if !valid {
		return false
	}

	// check after word has valid letters
	afterLetter := []rune(after)[0]
	valid = false
	for _, r := range validLetters {
		if r == afterLetter {
			valid = true
		}
	}
	if !valid {
		return false
	}

	validNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// check before word has valid numbers
	beforeNumber := string([]rune(before)[1])
	beforeNumberInt, err := strconv.Atoi(beforeNumber)
	if err != nil {
		panic(err)
	}
	valid = false
	for _, n := range validNumbers {
		if n == beforeNumberInt {
			valid = true
		}
	}
	if !valid {
		return false
	}

	// check after word has valid numbers
	afterNumber := string([]rune(after)[1])
	afterNumberInt, err := strconv.Atoi(afterNumber)
	if err != nil {
		panic(err)
	}
	valid = false
	for _, n := range validNumbers {
		if n == afterNumberInt {
			valid = true
		}
	}
	if !valid {
		return false
	}

	return true
}

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

func getPieceName(pieceNotation string) string {
	// parse piece notation, e.g. ○ P
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

	// e.g. white Pawn
	return colorNames[circle] + " " + pieceNames[piece]
}
