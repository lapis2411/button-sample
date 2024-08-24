package entity

import (
	"testing"

	"lapis2411/ruriGo/types"
)

func TestHitareaBoxIn(t *testing.T) {

	type args struct {
		position types.Position
	}
	rect := types.NewRectangle(100, 200, types.Position{X: 50, Y: 80})
	basicBoxHitarea := NewBoxHitarea(rect, 0)
	inclinedBoxHitarea := NewBoxHitarea(rect, 30)
	tests := []struct {
		name         string
		usingBoxArea *BoxHitarea
		args         args
		want         bool
	}{
		{
			name:         "position in the box area",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 60, Y: 100}},
			want:         true,
		}, {
			name:         "position in the box area edge(min)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 0, Y: -20}},
			want:         true,
		}, {
			name:         "position in the box area edge(max)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 99, Y: 179}},
			want:         true,
		}, {
			name:         "position of x is out of box area edge(minus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: -100, Y: 100}},
			want:         false,
		}, {
			name:         "position of x is out of box area edge(plus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 150, Y: 100}},
			want:         false,
		}, {
			name:         "position of y is out of box area edge(minus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 60, Y: -200}},
			want:         false,
		}, {
			name:         "position of y is out of box area edge(plus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 60, Y: 1200}},
			want:         false,
		}, {
			name:         "position in the inclined box area",
			usingBoxArea: inclinedBoxHitarea,
			args:         args{position: types.Position{X: 80, Y: 108}},
			want:         true,
		}, {
			name:         "position in the inclined box area edge(min)",
			usingBoxArea: inclinedBoxHitarea,
			args:         args{position: types.Position{X: 55, Y: -30}},
			want:         true,
		}, {
			name:         "position in the inclined box area edge(max)",
			usingBoxArea: inclinedBoxHitarea,
			args:         args{position: types.Position{X: 44, Y: 190}},
			want:         true,
		}, {
			name:         "position of x is out of inclined box area edge(minus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 127, Y: 145}},
			want:         false,
		}, {
			name:         "position of x is out of inclined box area edge(plus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 194, Y: -160}},
			want:         false,
		}, {
			name:         "position of y is out of inclined box area edge(minus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: -484, Y: 1064}},
			want:         false,
		}, {
			name:         "position of y is out of inclined box area edge(plus)",
			usingBoxArea: basicBoxHitarea,
			args:         args{position: types.Position{X: 601, Y: 1054}},
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ba := tt.usingBoxArea
			got := ba.Hit(tt.args.position)
			if got != tt.want {
				t.Errorf("%v: want %v, but %v", tt.name, tt.want, got)
			}
		})
	}

}
