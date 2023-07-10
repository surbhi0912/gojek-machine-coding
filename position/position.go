package position

type Position struct {
	x int
	y int
}

func New(xVal int, yVal int) Position {
	position := Position{
		x: xVal,
		y: yVal,
	}
	return position
}
