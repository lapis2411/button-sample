package resource

import (
	"bytes"
	_ "embed"
	"image/color"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/types"
)

//go:embed image/title.png
var backgroundImage []byte

type TitleImages struct {
	StartButton    *ebiten.Image
	ContinueButton *ebiten.Image
	ExitButton     *ebiten.Image
	NextButton     *ebiten.Image
	BackButton     *ebiten.Image
	MouseState     *ebiten.Image
	Background     *ebiten.Image
}

func LoadTitleImages() TitleImages {
	blue := color.RGBA{173, 216, 230, 255}
	bg, err := png.Decode(bytes.NewReader(backgroundImage))
	if err != nil {
		log.Fatal(err)
	}
	return TitleImages{
		StartButton:    ebiten.NewImageFromImage(createBoxImage("START", types.Size{Width: 100, Height: 30}, blue)),
		ContinueButton: ebiten.NewImageFromImage(createBoxImage("Continue", types.Size{Width: 100, Height: 30}, blue)),
		ExitButton:     ebiten.NewImageFromImage(createBoxImage("Exit", types.Size{Width: 100, Height: 30}, blue)),
		NextButton:     ebiten.NewImageFromImage(createBoxImage("Next", types.Size{Width: 100, Height: 30}, blue)),
		BackButton:     ebiten.NewImageFromImage(createBoxImage("Back", types.Size{Width: 100, Height: 30}, blue)),
		MouseState:     ebiten.NewImageFromImage(createBoxImage("LB", types.Size{Width: 50, Height: 50}, blue)),
		Background:     ebiten.NewImageFromImage(bg),
	}
}
