package main

import (
	"fmt"
	"gojek-machine-coding/player"
)

func main() {
	fmt.Println("Hello world")
	player.Player1 = player.NewPlayer(5, 5, 10)
	player.Player2 = player.NewPlayer(5, 5, 10)

	positionPlayer1 := [5][2]int{{1, 1}, {2, 0}, {2, 3}, {3, 4}}
	player.Player1.PlayerBattleSpace.InitialiseShip(positionPlayer1)
}
