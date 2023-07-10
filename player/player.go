package player

import (
	"gojek-machine-coding/grid"
	"gojek-machine-coding/position"
)

type Player struct {
	TotalMissiles     int
	RemainingMissiles int
	PlayerBattleSpace grid.Grid
	Positions         []position.Position
}

//func New(totMissiles int) *Player {
//	player := Player{
//		TotalMissiles:     totMissiles,
//		RemainingMissiles: totMissiles,
//	}
//}
