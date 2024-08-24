package types

import (
	"strconv"
)

type Position struct {
	X, Y int
}

func (p Position) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func (p Position) Add(q Position) Position {
	return Position{p.X + q.X, p.Y + q.Y}
}

func (p Position) ToDrawPosition(size Size) Position {
	dx := size.Width / 2
	dy := size.Height / 2
	return Position{X: p.X - int(dx), Y: p.Y - int(dy)}
}
