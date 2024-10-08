package entity

import (
	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/types"
)

type GameObject interface {
	Position() types.Position
	IsEnable() bool
}

type Hitarea interface {
	Hit(pos types.Position) bool
	SetReferencePosition(pos types.Position)
	IsEnable() bool
	Activate()
	Deactivate()
}

type Object struct {
	position types.Position
	Image    *ebiten.Image
	enable   bool
}

func NewObject(position types.Position, image *ebiten.Image) *Object {
	return &Object{
		position: position,
		Image:    image,
		enable:   true,
	}
}

func (o *Object) SetPosition(position types.Position) {
	o.position = position
}

func (o *Object) Position() types.Position {
	return o.position
}

func (o *Object) IsEnable() bool {
	return o.enable
}

func (o *Object) Enable() {
	o.enable = true
}

func (o *Object) Disable() {
	o.enable = false
}

type InteractiveObject struct {
	*Object
	hitarea Hitarea
}

func (iobj *InteractiveObject) SetPosition(position types.Position) {
	iobj.hitarea.SetReferencePosition(position)
	iobj.Object.SetPosition(position)
}
