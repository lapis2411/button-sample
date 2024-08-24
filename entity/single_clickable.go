package entity

import (
	"lapis2411/button-sample/types"
)

type Clickable interface {
	UpdateStatus(point types.Position, clicked bool) error
	IsClicked() bool
	UnFocus()
}

type SingleClickableGroup struct {
	clickables        []Clickable
	clickedObjectPrev bool
}

func NewSingleClickableGroup(c []Clickable) SingleClickableGroup {
	return SingleClickableGroup{
		clickables:        c,
		clickedObjectPrev: false,
	}
}

func (cg *SingleClickableGroup) Click(point types.Position, clicked bool) {
	cg.clickedObjectPrev = cg.clickedObjectPrev && clicked // クリック解除でロックも解除
	for _, obj := range cg.clickables {
		// 他のボタンがクリックされ続けている場合何もしない
		clickContinue := cg.clickedObjectPrev && clicked
		if clickContinue && !obj.IsClicked() {
			obj.UnFocus()
			continue
		}
		obj.UpdateStatus(point, clicked)
		if obj.IsClicked() {
			cg.clickedObjectPrev = true
		}
	}
}
