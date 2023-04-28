package main

import (
	"fmt"
)

func main() {
	firstPlayer := Player{"X", "Player 1"}
	secondPlayer := Player{"O", "Player 2"}

	fmt.Println(fmt.Sprintf("Player 1 has %s", firstPlayer.symbol))
	fmt.Println(fmt.Sprintf("Player 2 has %s", secondPlayer.symbol))

	var board Board
	board.printBoard()
}

type Player struct {
	symbol string
	name   string
}

type Board struct {
	tiles [3][3]string
}

func (b Board) printBoard() {
	fmt.Println("+-+-+-+")
	for i := range b.tiles {
		for j := range b.tiles[i] {
			symbol := b.tiles[i][j]
			if symbol == "" {
				fmt.Print("|", " ")
			} else {
				fmt.Print("|", symbol)
			}
		}
		fmt.Println("|\n+-+-+-+")
	}
}
