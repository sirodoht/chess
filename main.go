package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// initialize board
	board := Board{}
	board.Init()

	// main game loop
	for {
		board.Render()

		// read from terminal
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter next move: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command = strings.TrimSpace(command)

		move, err := NewMove(command)
		if err != nil {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		// execute move
		pieceName, afterCommand := board.Execute(move)
		fmt.Printf("\nMOVE: %s moved to %s\n", pieceName, afterCommand)

		// check for exit
		if command == "exit" {
			break
		}
	}
}
