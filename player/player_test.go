package player

import "testing"

type playerTest struct {
	GridSize          int
	TotalShips        int
	TotalMissiles     int
	RemainingMissiles int
	want              *Player
}

func TestNewPlayer(t *testing.T) {
	test := playerTest{
		GridSize:          5,
		TotalShips:        5,
		TotalMissiles:     10,
		RemainingMissiles: 10,
		want: &Player{},
	}
}
