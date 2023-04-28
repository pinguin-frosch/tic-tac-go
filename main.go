package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tic Tac Go!")
	var game Game
	game.setup()
	game.play()
}
