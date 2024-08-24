package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/entity"
	"lapis2411/button-sample/types"
)

const UI_CLICK_EFFECT_COUNT = 5
const EFFECT_FOCUS_EXPAND = 1.05
const EFFECT_CLICK_EXPAND = 1.1
const EFFECT_CLICK_SHRINK = 0.9

func DrawButton(screen *ebiten.Image, button entity.Button, image *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := types.Size{Width: 1, Height: 1}
	if button.IsClicked() {
		scale = getScaleByCount(button.ClickingFrameCount())
	} else if button.IsFocused() {
		scale = types.Size{Width: EFFECT_FOCUS_EXPAND, Height: EFFECT_FOCUS_EXPAND}
	}
	op.GeoM.Scale(scale.Width, scale.Height)

	b := image.Bounds()
	size := types.Size{
		Width:  float64(b.Dx()) * scale.Height,
		Height: float64(b.Dy()) * scale.Height,
	}
	pos := button.Position().ToDrawPosition(size)
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	op.ColorScale.ScaleAlpha(1)
	screen.DrawImage(image, op)
}

func getScaleByCount(count int) types.Size {
	th := UI_CLICK_EFFECT_COUNT / 2
	if count > th {
		return types.Size{Width: EFFECT_CLICK_SHRINK, Height: EFFECT_CLICK_SHRINK}
	}
	return types.Size{Width: EFFECT_CLICK_EXPAND, Height: EFFECT_CLICK_EXPAND}
}
