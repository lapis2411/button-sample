package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/entity"
	"lapis2411/button-sample/types"
)

type Title3 struct {
	selector TitleSelection
}

func NewTitle3() (*Title3, error) {
	return &Title3{
		selector: None,
	}, nil
}

func (t *Title3) Initialize() (*data.Title, error) {
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
		data.WithText("This is Title3"),
		data.WithButtons(btns))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize title: %w", err)
	}

	d.TitleIndex = 3

	return &d, err
}

func (t *Title3) Update(data *data.Title) error {
	mx, my := ebiten.CursorPosition()
	cursorPosition := types.Position{X: mx, Y: my}
	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	justClicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
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

func (t *Title3) Selector() TitleSelection {
	return t.selector
}
