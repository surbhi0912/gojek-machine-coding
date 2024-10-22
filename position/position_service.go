package position

func NewPosition(xVal int, yVal int) Position {
	position := Position{
		x: xVal,
		y: yVal,
	}
	return position
}
