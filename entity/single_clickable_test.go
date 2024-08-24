package entity

import (
	"reflect"
	"testing"

	"lapis2411/ruriGo/types"
)

func TestSingleClickableGroupClick(t *testing.T) {
	const numberOfButtons = 3
	type args struct {
		position types.Position
		click    bool
	}
	type wants struct {
		clickNumber   int
		buttonClicked []bool
		buttonFocused []bool
		clickedPrev   bool
	}

	clickableObjs := []Clickable{}
	btns := make([]*Button, numberOfButtons)
	clickNumber := -1
	// initialize button 0,1,2
	for i, y := range []int{100, 200, 300} {
		b := NewRectangleButtonWithElement(
			100,
			50,
			types.Position{X: 100, Y: y},
			func(b Button) error {
				if b.FocusAndJustReleased() {
					clickNumber = i
				}
				return nil
			})
		btns[i] = b
		clickableObjs = append(clickableObjs, b)
	}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "focus button 1",
			args: args{position: types.Position{X: 110, Y: 210}, click: false},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, false, false},
				buttonFocused: []bool{false, true, false},
				clickedPrev:   false,
			},
		}, {
			name: "click button 1",
			args: args{position: types.Position{X: 110, Y: 210}, click: true},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, true, false},
				buttonFocused: []bool{false, true, false},
				clickedPrev:   true,
			},
		}, {
			name: "clicked button 1 prev, but unfocused with clicking",
			args: args{position: types.Position{X: 1100, Y: 2100}, click: true},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, true, false},
				buttonFocused: []bool{false, false, false},
				clickedPrev:   true,
			},
		}, {
			name: "clicked button 1 prev continue, but focus button 2",
			args: args{position: types.Position{X: 110, Y: 310}, click: true},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, true, false},
				buttonFocused: []bool{false, false, false},
				clickedPrev:   true,
			},
		}, {
			name: "clicked button 1 prev continue, but focus button 1",
			args: args{position: types.Position{X: 110, Y: 310}, click: true},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, true, false},
				buttonFocused: []bool{false, false, false},
				clickedPrev:   true,
			},
		}, {
			name: "button 1 clicked prev, focus button 1 now, and click is relieved",
			args: args{position: types.Position{X: 110, Y: 310}, click: false},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, false, false},
				buttonFocused: []bool{false, false, true},
				clickedPrev:   false,
			},
		}, {
			name: "click button 1 (for prepare)",
			args: args{position: types.Position{X: 110, Y: 210}, click: true},
			wants: wants{
				clickNumber:   -1,
				buttonClicked: []bool{false, true, false},
				buttonFocused: []bool{false, true, false},
				clickedPrev:   true,
			},
		}, {
			name: "click button 1 prev, and click is relieved now",
			args: args{position: types.Position{X: 110, Y: 210}, click: false},
			wants: wants{
				clickNumber:   1,
				buttonClicked: []bool{false, false, false},
				buttonFocused: []bool{false, true, false},
				clickedPrev:   false,
			},
		},
	}
	singleClickables := NewSingleClickableGroup(clickableObjs)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			singleClickables.Click(tt.args.position, tt.args.click)
			got := wants{
				clickNumber:   clickNumber,
				buttonClicked: []bool{btns[0].IsClicked(), btns[1].IsClicked(), btns[2].IsClicked()},
				buttonFocused: []bool{btns[0].IsFocused(), btns[1].IsFocused(), btns[2].IsFocused()},
				clickedPrev:   singleClickables.clickedObjectPrev,
			}
			if !reflect.DeepEqual(got, tt.wants) {
				t.Errorf("%v: want %v, but %v", tt.name, tt.wants, got)
			}
		})
	}

}
