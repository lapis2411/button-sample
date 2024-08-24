package entity

import (
	"errors"
	"math"

	"lapis2411/button-sample/types"
)

var _ Hitarea = (*BoxHitarea)(nil)

type BoxHitarea struct {
	rectangle types.Rectangle
	radian    float64
	enable    bool
}

func NewBoxHitarea(rect types.Rectangle, angle float64) *BoxHitarea {
	rad := float64(angle) * math.Pi / 180
	return &BoxHitarea{
		rectangle: rect,
		radian:    rad,
		enable:    true,
	}
}

func (b *BoxHitarea) Clone() *BoxHitarea {
	return &BoxHitarea{
		rectangle: b.rectangle,
		radian:    b.radian,
		enable:    b.enable,
	}
}

func (b *BoxHitarea) Hit(pos types.Position) bool {
	c := b.rectangle.Center()
	cos := math.Cos(-1 * b.radian)
	sin := math.Sin(-1 * b.radian)
	x := float64(pos.X-c.X)*cos - float64(pos.Y-c.Y)*sin + float64(c.X)
	y := float64(pos.Y-c.Y)*cos + float64(pos.X-c.X)*sin + float64(c.Y)
	rpos := types.Position{X: int(x), Y: int(y)}
	return b.rectangle.In(rpos)
}

func (b *BoxHitarea) SetReferencePosition(position types.Position) {
	w := b.rectangle.Width()
	h := b.rectangle.Height()
	b.rectangle = types.NewRectangle(w, h, position)
}

func (b *BoxHitarea) SetAngle(angle float64) error {
	if angle < 0 || 90 <= angle {
		return errors.New("angle error. angle must be 0 <= angle < 90")
	}
	b.radian = float64(angle) * math.Pi / 180
	return nil
}

func (b *BoxHitarea) IsEnable() bool {
	return b.enable
}

func (b *BoxHitarea) Activate() {
	b.enable = true
}

func (b *BoxHitarea) Deactivate() {
	b.enable = false
}
