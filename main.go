package main

import (
	"fmt"
)

func main() {
	firstPlayer := Player{"X", "Player 1"}
	secondPlayer := Player{"O", "Player 2"}

	fmt.Println("Tic Tac Go!")

	players := [2]Player{firstPlayer, secondPlayer}
	var board Board

	for i := 0; true; i = (i + 1) % 2 {
		player := players[i]
		board.printBoard()

		fmt.Println(fmt.Sprintf("%s (%s)", player.name, player.symbol))
		fmt.Print("Enter row and column tile: ")

		var row, col int

		for {
			fmt.Scanf("%d %d", &row, &col)
			if board.validTile(row, col) {
				break
			} else {
				fmt.Print("Invalid tile location! Try again: ")
			}
		}

		board.usedTiles++
		board.tiles[row][col] = player.symbol

		state := board.checkWin(row, col, player.symbol)
		if state == 1 {
			board.printBoard()
			fmt.Println(fmt.Sprintf("%s (%s) wins!", player.name, player.symbol))
			break
		} else if state == -1 {
			board.printBoard()
			fmt.Println("It's a tie!")
			break
		}
	}
}

type Player struct {
	symbol string
	name   string
}

type Board struct {
	tiles     [3][3]string
	usedTiles uint
}

func (b Board) validTile(row int, col int) bool {
	if row >= len(b.tiles) || col >= len(b.tiles[0]) {
		return false
	}

	if b.tiles[row][col] != "" {
		return false
	}

	return true
}

func (b Board) checkWin(row int, col int, symbol string) int {
	directions := map[string]int{"v": 0, "h": 0, "r": 0, "l": 0}
	for i := 0; i < 3; i++ {
		if b.tiles[i][col] != symbol {
			break
		}
		directions["v"]++
	}
	for j := 0; j < 3; j++ {
		if b.tiles[row][j] != symbol {
			break
		}
		directions["h"]++
	}
	for k := 0; k < 3; k++ {
		if b.tiles[k][k] != symbol {
			break
		}
		directions["l"]++
	}
	for l := 0; l < 3; l++ {
		if b.tiles[2-l][l] != symbol {
			break
		}
		directions["r"]++
	}

	for k := range directions {
		if directions[k] == 3 {
			return 1
		}
	}

	if b.usedTiles < 9 {
		return 0
	} else {
		return -1
	}
}

func (b Board) printBoard() {
	fmt.Println("\n+-+-+-+")
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
	fmt.Println()
}
