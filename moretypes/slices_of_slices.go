package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{ // HL
		[]string{"_", "_", "_"}, // HL
		[]string{"_", "_", "_"}, // HL
		[]string{"_", "_", "_"}, // HL
	} // HL

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// END OMIT
