package player

import (
	"gojek-machine-coding/grid"
	"gojek-machine-coding/position"
)

type Player struct {
	GridSize          int
	TotalShips        int
	TotalMissiles     int
	RemainingMissiles int
	PlayerBattleSpace grid.Grid
	ShipPositions     []position.Position
	MissilePositions  []position.Position
}
