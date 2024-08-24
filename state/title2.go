package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/entity"
	"lapis2411/button-sample/types"
)

type Title2 struct {
	selector   TitleSelection
	clickables entity.SingleClickableGroup
}

func NewTitle2() (*Title2, error) {
	return &Title2{
		selector:   None,
		clickables: entity.SingleClickableGroup{},
	}, nil
}

func (t *Title2) Initialize() (*data.Title, error) {
	clickableObjs := []entity.Clickable{}
	btns := make([]*entity.Button, NumberButtons)
	for i, y := range []int{StartButtonPositionY, ContinueButtonPositionY, ExitButtonPositionY, NextButtonPositionY, BackButtonPositionY} {
		b := entity.NewRectangleButtonWithElement(
			ButtonWidth,
			ButtonHeight,
			types.Position{X: ButtonPositionX, Y: y},
			func(button entity.Button) error {
				if button.JustReleased() {
					t.selector = TitleSelection(i + 1)
				}
				return nil
			})
		btns[i] = b
		clickableObjs = append(clickableObjs, b)
	}

	d, err := data.NewTitle(
		data.WithText("This is Title2"),
		data.WithButtons(btns))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize title: %w", err)
	}

	d.TitleIndex = 2
	t.clickables = entity.NewSingleClickableGroup(clickableObjs)

	return &d, err
}

func (t *Title2) Update(data *data.Title) error {
	mx, my := ebiten.CursorPosition()
	cursorPosition := types.Position{X: mx, Y: my}
	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	t.clickables.Click(cursorPosition, mouseClicked)
	if t.selector == TitleStart {
		data.ButtonText = "START Button Clicked"
	} else if t.selector == TitleContinue {
		data.ButtonText = "ContinueButton Button Clicked"
	} else if t.selector == TitleExit {
		data.ButtonText = "ExitButton Button Clicked"
	}
	return nil
}

func (t *Title2) Selector() TitleSelection {
	return t.selector
}
