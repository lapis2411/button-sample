package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/entity"
	"lapis2411/button-sample/types"
)

const (
	NumberButtons = 5

	ButtonWidth  = 100
	ButtonHeight = 30

	ButtonPositionX         = 400
	StartButtonPositionY    = 50
	ContinueButtonPositionY = 100
	ExitButtonPositionY     = 150
	NextButtonPositionY     = 200
	BackButtonPositionY     = 250
)

type TitleSelection int
type Title struct {
	selector TitleSelection
}

const (
	None TitleSelection = iota
	TitleStart
	TitleContinue
	TitleExit
	TitleNext
	TitleBack
)

func NewTitle() (*Title, error) {
	return &Title{
		selector: None,
	}, nil
}

func (t *Title) Initialize() (*data.Title, error) {
	btns := make([]*entity.Button, NumberButtons)
	for i, y := range []int{StartButtonPositionY, ContinueButtonPositionY, ExitButtonPositionY, NextButtonPositionY, BackButtonPositionY} {
		b := entity.NewRectangleButton(
			types.NewRectangle(ButtonWidth, ButtonHeight, types.Position{X: ButtonPositionX, Y: y}),
			entity.WithButtonEvent(
				func(button entity.Button) error {
					if button.IsClicked() {
						t.selector = TitleSelection(i + 1)
					}
					return nil
				}),
		)
		btns[i] = b
	}

	d, err := data.NewTitle(
		data.WithText("This is Title"),
		data.WithButtons(btns))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize title: %w", err)
	}

	d.TitleIndex = 1
	return &d, err
}

func (t *Title) Update(data *data.Title) error {
	mx, my := ebiten.CursorPosition()
	cursorPosition := types.Position{X: mx, Y: my}
	mouseClicked := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
	justClicked := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
	for _, b := range data.Buttons() {
		b.UpdateStatus(cursorPosition, mouseClicked, justClicked)
		b.UnFocus()
	}
	if t.selector == TitleStart {
		data.ButtonText = "START Button Clicked"
	} else if t.selector == TitleContinue {
		data.ButtonText = "ContinueButton Button Clicked"
	} else if t.selector == TitleExit {
		data.ButtonText = "ExitButton Button Clicked"
	}
	return nil
}

func (t *Title) Selector() TitleSelection {
	return t.selector
}
