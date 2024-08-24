package data

import (
	"lapis2411/button-sample/entity"
)

type Title struct {
	ButtonText string
	TitleIndex int
	text       string
	buttons    []*entity.Button
}

type TitleOption func(*Title)
type TitleButton int

const (
	StartButton TitleButton = iota
	ContinueButton
	ExitButton
	NextButton
	BackButton
)

func NewTitle(opts ...TitleOption) (Title, error) {
	t := Title{}

	for _, option := range opts {
		option(&t)
	}
	return t, nil
}

func WithText(text string) TitleOption {
	return func(t *Title) {
		t.text = text
	}
}

func WithButtons(buttons []*entity.Button) TitleOption {
	return func(t *Title) {
		t.buttons = buttons
	}
}

func (t *Title) Text() string {
	return t.text
}

func (t *Title) Buttons() []*entity.Button {
	return t.buttons
}
