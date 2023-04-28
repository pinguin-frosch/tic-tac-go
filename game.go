package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Game struct {
	players     [2]Player
	playerIndex uint
	board       Board
	games       uint
	gameIndex   uint
}

func (g *Game) play() {
	for {
		g.playRound()
		g.printStats()
		g.board.resetBoard()
		g.playerIndex = (g.playerIndex + 1) % 2
		g.gameIndex++
		if g.gameIndex >= g.games {
			fmt.Println("Game over.")
			if g.players[0].wins > g.players[1].wins {
				fmt.Printf("%s wins!\n", g.players[0].name)
			} else if g.players[0].wins < g.players[1].wins {
				fmt.Printf("%s wins!\n", g.players[1].name)
			} else {
				fmt.Println("It's a tie!")
				var response string
				fmt.Printf("Do you want to play another round? (yes/no): ")
				fmt.Scanln(&response)
				if strings.ToLower(response) == "yes" {
					fmt.Println("\nGet ready for another round!")
					fmt.Println()
					g.games++
					continue
				}
			}
			break
		}
	}
}

func (g *Game) printStats() {
	fmt.Println("\nTotal stats:")
	fmt.Printf("%s %d - %d %s\n", g.players[0].name, g.players[0].wins, g.players[1].wins, g.players[1].name)
	fmt.Scanf("Press to continue playing...")
}

func (g *Game) playRound() {
	fmt.Printf("%s will start round %d!", g.players[g.playerIndex].name, g.gameIndex+1)

	for i := g.playerIndex; true; i = (i + 1) % 2 {
		player := g.players[i]
		g.board.printBoard()

		fmt.Printf("%s's turn (%s)\n", player.name, player.symbol)
		fmt.Print("Enter row and column tile: ")

		var row, col int

		for {
			fmt.Scanf("%d %d", &row, &col)
			if g.board.validTile(row, col) {
				break
			} else {
				fmt.Print("Invalid tile location! Try again: ")
			}
		}

		g.board.writeSymbol(row, col, player.symbol)

		state := g.board.checkWin(row, col, player.symbol)
		if state == 1 {
			g.board.printBoard()
			g.players[i].wins++
			fmt.Printf("%s wins!\n", player.name)
			break
		} else if state == -1 {
			g.board.printBoard()
			fmt.Println("It's a tie!")
			break
		}
	}
}

func (g *Game) setup() {
	g.players[0].symbol = "X"
	g.players[1].symbol = "O"

	fmt.Print("\nHow many rounds will there be? ")
	_, err := fmt.Scanf("%d", &g.games)
	if err != nil {
		fmt.Println("Not a valid numbers! Using 3 by default.")
		g.games = 3
	}

	fmt.Print("\nWho will be X? ")
	fmt.Scanln(&g.players[0].name)
	fmt.Print("Who will be O? ")
	fmt.Scanln(&g.players[1].name)
	fmt.Println()

	g.playerIndex = uint(rand.Intn(2))
}

type Player struct {
	symbol string
	name   string
	wins   uint
}
