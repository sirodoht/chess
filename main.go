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
		move, isValid, msg := NewMove(board, turn, command)
		if !isValid {
			fmt.Printf("\n%s\n", msg)
			continue
		}

		// execute move
		board.Execute(move)

		// show status message
		fmt.Printf("\n%s\n", msg)

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
