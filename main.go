package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var board Board
	board.printBoard()
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
