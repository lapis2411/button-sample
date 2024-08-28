package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/entity"
	"lapis2411/button-sample/types"
)

type Title4 struct {
	selector TitleSelection
}

func NewTitle4() (*Title4, error) {
	return &Title4{
		selector: None,
	}, nil
}

func (t *Title4) Initialize() (*data.Title, error) {
	btns := make([]*entity.Button, NumberButtons)
	for i, y := range []int{StartButtonPositionY, ContinueButtonPositionY, ExitButtonPositionY, NextButtonPositionY, BackButtonPositionY} {
		b := entity.NewRectangleButton(
			types.NewRectangle(ButtonWidth, ButtonHeight, types.Position{X: ButtonPositionX, Y: y}),
			entity.WithButtonEvent(
				func(button entity.Button) error {
					if button.FocusAndJustReleased() {
						t.selector = TitleSelection(i + 1)
					}
					return nil
				}),
		)
		btns[i] = b
	}

	d, err := data.NewTitle(
		data.WithText("This is Title4"),
		data.WithButtons(btns))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize title: %w", err)
	}

	d.TitleIndex = 4

	return &d, err
}

func (t *Title4) Update(data *data.Title) error {
	mx, my := ebiten.CursorPosition()
	cursorPosition := types.Position{X: mx, Y: my}
	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	justClicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	for _, b := range data.Buttons() {
		b.UpdateStatus(cursorPosition, mouseClicked, justClicked)
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

func (t *Title4) Selector() TitleSelection {
	return t.selector
}
