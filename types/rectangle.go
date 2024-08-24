package types

import (
	"fmt"
	"log/slog"
)

type Rectangle struct {
	Min, Max Position
}

func NewRectangle(width, height int, center Position) Rectangle {
	if width%2 != 0 || height%2 != 0 {
		slog.Info(fmt.Sprintf("The width or height is not a multiple of 2. width: %v, height: %v", width, height))
	}
	dw := width / 2
	dh := height / 2
	mini := Position{
		X: center.X - dw,
		Y: center.Y - dh,
	}
	maxi := Position{
		X: center.X + dw,
		Y: center.Y + dh,
	}
	return Rectangle{Min: mini, Max: maxi}
}

func (r Rectangle) Width() int {
	return r.Max.X - r.Min.X
}

func (r Rectangle) Height() int {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) Center() Position {
	w := r.Max.X - r.Min.X
	h := r.Max.Y - r.Min.Y
	mini := r.Min
	hf := Position{X: w / 2, Y: h / 2}
	return mini.Add(hf)
}

func (r Rectangle) In(pos Position) bool {
	return r.Min.X <= pos.X && pos.X < r.Max.X &&
		r.Min.Y <= pos.Y && pos.Y < r.Max.Y
}
