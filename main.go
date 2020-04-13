package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// initialize game
	board := Board{}
	board.Init()
	turn := WHITE

	// main game loop
	for {
		board.Render()

		// read from stdin
		reader := bufio.NewReader(os.Stdin)
		turnName := GetTeamName(turn, UPPER)
		fmt.Printf("%s plays. Enter next move: ", turnName)
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command = strings.TrimSpace(command)

		// create move
		move, err := NewMove(board, turn, command)
		if err != nil {
			fmt.Println("\nInvalid move. Please try again.")
			continue
		}

		// execute move
		pieceName, afterCommand := board.Execute(move)
		fmt.Printf("\nMOVE: %s moved to %s\n", pieceName, afterCommand)

		// check for exit
		if command == "exit" {
			break
		}

		// change turns
		if turn == WHITE {
			turn = BLACK
		} else {
			turn = WHITE
		}
	}
}
