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
	board.Render()

	// main game loop
	for {
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
		move, isValid, messages, isEndgame := NewMove(board, turn, command)

		// check move validity
		if !isValid {
			board.Render()
			if len(messages) > 0 {
				fmt.Printf("%s\n", messages[0])
			}
			continue
		}

		// execute move
		board.Execute(move)

		// render new board
		board.Render()

		// show status message, start from i=1
		for i := 1; i < len(messages); i++ {
			fmt.Printf("%s\n", messages[i])
		}

		if isEndgame {
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
