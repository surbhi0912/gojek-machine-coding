package player

import (
	"gojek-machine-coding/grid"
	"gojek-machine-coding/position"
)

func NewPlayer(gridSize, numShips, numMissiles int) *Player {
	player := &Player{
		GridSize:          gridSize,
		TotalShips:        numShips,
		TotalMissiles:     numMissiles,
		RemainingMissiles: numMissiles,
		ShipPositions:     make([]position.Position, numShips),
		MissilePositions:  make([]position.Position, numMissiles),
	}
	player.PlayerBattleSpace = grid.New(gridSize)
	return player
}

func (p *Player) InitialiseShipPositions(pos position.Position) {

}
