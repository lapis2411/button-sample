package entity

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/types"
)

type Button struct {
	*InteractiveObject
	focused    bool
	clicked    bool
	preClicked bool
	event      ButtonEvent
}

type ButtonOption func(*Button)
type ButtonEvent func(Button) error

func NewButton(position types.Position, hitarea Hitarea, opts ...ButtonOption) Button {
	obj := InteractiveObject{
		Object:  NewObject(position, nil),
		hitarea: hitarea,
	}

	b := Button{
		InteractiveObject: &obj,
		focused:           false,
		clicked:           false,
		preClicked:        false,
	}
	for _, option := range opts {
		option(&b)
	}
	return b
}

func NewRectangleButton(rect types.Rectangle, opts ...ButtonOption) *Button {
	ha := NewBoxHitarea(rect, 0)
	b := NewButton(rect.Center(), ha, opts...)
	return &b
}

func WithButtonEvent(event ButtonEvent) ButtonOption {
	return func(b *Button) {
		b.event = event
	}
}

func WithImage(image *ebiten.Image) ButtonOption {
	return func(b *Button) {
		b.Object.Image = image
	}
}

func (b *Button) SetButtonEvent(event ButtonEvent) {
	b.event = event
}

func (b *Button) UpdateStatus(point types.Position, clicked bool, justClicked bool) error {
	if !b.IsEnable() {
		return nil
	}
	b.focused = b.InteractiveObject.hitarea.Hit(point)
	b.preClicked = b.clicked

	//　クリックの状態が変わるのはフォーカスされてクリックした時かクリックを離した時
	// focusを外してクリックしている状態の時はときに選択を変えない
	if !clicked {
		b.clicked = false
	} else if b.focused && justClicked {
		b.clicked = true
	}
	if b.event != nil {
		return b.event(*b)
	}
	return nil
}

func (b *Button) UnFocus() {
	b.focused = false
}

func (b *Button) IsFocused() bool {
	return b.focused
}

func (b *Button) IsClicked() bool {
	return b.clicked
}

func (b *Button) JustReleased() bool {
	return b.preClicked && !b.clicked
}

func (b *Button) FocusAndJustReleased() bool {
	return b.focused && b.preClicked && !b.clicked
}

func (b *Button) DebugPrint() string {
	return fmt.Sprintf("position: %v, focused: %v, clicked: %v", b.position, b.focused, b.clicked)
}
