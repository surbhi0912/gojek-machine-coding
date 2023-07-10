package grid

func New(gridSize, numShips int) Grid {
	grid := Grid{
		GridSize:    gridSize,
		NumShips:    numShips,
		BattleSpace: make([][]byte, gridSize),
	}
	for i := 0; i < gridSize; i++ {
		grid.BattleSpace[i] = make([]byte, gridSize)
	}
	return grid
}
