package entity

import (
	"reflect"
	"testing"

	"lapis2411/button-sample/types"
)

func TestUpdateStatus(t *testing.T) {

	center := types.Position{X: 10, Y: 10}
	rect := types.NewRectangle(20, 20, center)
	ha := NewBoxHitarea(rect, 0)
	btn := NewButton(center, ha)
	btn2 := NewButton(center, ha)
	type args struct {
		point       types.Position
		clicked     bool
		jsutClicked bool
	}
	type wants struct {
		focused bool
		clicked bool
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{
			name: "Not focused Not Clicked",
			args: args{point: types.Position{X: 60, Y: 100}, clicked: false, jsutClicked: false},
			want: wants{focused: false, clicked: false},
		}, {
			name: "Focused Not Clicked",
			args: args{point: types.Position{X: 10, Y: 15}, clicked: false, jsutClicked: false},
			want: wants{focused: true, clicked: false},
		}, {
			name: "Not Focused, Clicked",
			args: args{point: types.Position{X: 100, Y: 15}, clicked: true, jsutClicked: true},
			want: wants{focused: false, clicked: false},
		}, {
			name: "Focused and Clicked",
			args: args{point: types.Position{X: 5, Y: 15}, clicked: true, jsutClicked: true},
			want: wants{focused: true, clicked: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := btn
			b.UpdateStatus(tt.args.point, tt.args.clicked, tt.args.jsutClicked)
			got := wants{
				focused: b.focused,
				clicked: b.clicked,
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v: want %v, but %v", tt.name, tt.want, got)
			}
		})
	}

	t.Run("check unfocuse with clicking", func(t *testing.T) {
		t.Parallel()
		btn2.UpdateStatus(types.Position{X: 5, Y: 15}, true, true)
		// clickしたままFocusを外してもクリックはtrueのまま
		btn2.UpdateStatus(types.Position{X: 1000, Y: 0}, true, false)
		if !btn2.clicked {
			t.Errorf("when focus is out being clicked, must be clicked")
		}
		if btn2.focused {
			t.Errorf("focus is out")
		}

		raised := false
		btn2.SetButtonEvent(func(b Button) error {
			if b.FocusAndJustReleased() {
				raised = true
			}
			return nil
		})

		btn2.UpdateStatus(types.Position{X: 0, Y: 0}, false, false)
		if !raised {
			t.Errorf("clicked event want to be raised, but not raised")
		}
		if btn2.clicked {
			t.Errorf("not clicked, so must be not clicked")
		}
		if !btn2.focused {
			t.Errorf("still be focused")
		}

		// キャンセル(focusを外してクリックを止める)パターン
		raised = false
		btn2.UpdateStatus(types.Position{X: 0, Y: 0}, true, true)
		for i := 0; i < 4; i++ {
			btn2.UpdateStatus(types.Position{X: 0, Y: 0}, true, false)
		}
		btn2.UpdateStatus(types.Position{X: -50, Y: 0}, false, false)
		if raised {
			t.Errorf("event must not be raised when focus out and unclick")
		}
		if btn2.clicked {
			t.Errorf("already unclicked")
		}

	})
}
