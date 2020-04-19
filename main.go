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
		fmt.Printf("%s plays. Enter next %s move: ", turnName, GetTeamName(turn, SYMBOL))
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command = strings.TrimSpace(command)

		// check for exit
		if command == "exit" || command == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		// check for resignation
		if command == "resigns" {
			winner := WHITE
			if turn == WHITE {
				winner = BLACK
			}
			fmt.Printf("RESIGNATION: %s wins!\n", GetTeamName(winner, LOWER))
			break
		}

		// create move
		move, isValid, msg, isEndgame := NewMove(board, turn, command)

		// check move validity
		if !isValid {
			fmt.Printf("\n%s\n", msg)
			continue
		}

		// execute move
		board.Execute(move)

		// show status message
		fmt.Printf("\n%s\n", msg)

		if isEndgame {
			board.Render()
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
