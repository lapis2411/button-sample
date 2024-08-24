package drawer

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/resource"
	"lapis2411/button-sample/types"
)

type Title struct {
	images resource.TitleImages
}

func NewTitle() (Title, error) {
	return Title{
		images: resource.LoadTitleImages(),
	}, nil
}

func (t Title) Update(s *ebiten.Image, dt data.Title) error {
	t.drawBackground(s, dt.TitleIndex)
	btns := dt.Buttons()
	DrawButton(s, *btns[data.StartButton], t.images.StartButton)
	DrawButton(s, *btns[data.ContinueButton], t.images.ContinueButton)
	DrawButton(s, *btns[data.ExitButton], t.images.ExitButton)
	DrawButton(s, *btns[data.NextButton], t.images.NextButton)
	DrawButton(s, *btns[data.BackButton], t.images.BackButton)
	txt := dt.Text() + "\n"
	for _, idx := range []data.TitleButton{data.StartButton, data.ContinueButton, data.ExitButton, data.NextButton, data.BackButton} {
		txt += fmt.Sprintf("%s\n", btns[idx].DebugPrint())
	}
	txt += dt.ButtonText
	ebitenutil.DebugPrint(s, txt)
	DrawMouseState(s, t.images.MouseState)

	return nil
}

func (t Title) drawBackground(s *ebiten.Image, titleID int) {
	bg := t.images.Background
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate()
	scale := types.Size{Width: 0.4, Height: 0.4}
	op.GeoM.Scale(scale.Width, scale.Height)
	// 画像を描画
	if titleID == 2 {
		op.ColorScale.Scale(1, 0.8, 0.8, 1)
	} else if titleID == 3 {
		op.ColorScale.Scale(1, 0.5, 0.5, 1)
	}

	s.DrawImage(bg, op)
}

func DrawMouseState(screen *ebiten.Image, image *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	scale := types.Size{Width: 1, Height: 1}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		scale = types.Size{Width: 0.8, Height: 0.8}
	}
	op.GeoM.Scale(scale.Width, scale.Height)

	op.GeoM.Translate(float64(50), float64(200))
	op.ColorScale.ScaleAlpha(1)
	screen.DrawImage(image, op)
}
