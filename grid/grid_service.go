package grid

import (
	"fmt"
)

func New(gridSize int) Grid {
	grid := Grid{
		BattleSpace: make([][]byte, gridSize),
	}
	for i := 0; i < gridSize; i++ {
		grid.BattleSpace[i] = make([]byte, gridSize)
	}
	return grid
}

func (g *Grid) InitialiseShip(pos [5][2]int) {
	for _, p := range pos {
		fmt.Println(p[0], p[1])
	}
}
