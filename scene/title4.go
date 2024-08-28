package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/drawer"
	"lapis2411/button-sample/state"
)

var _ Scene = (*Title4)(nil)

type Title4 struct {
	drawer drawer.Title
	state  *state.Title4
	data   *data.Title
}

func NewTitle4() (Title4, error) {
	dr, err := drawer.NewTitle()
	if err != nil {
		return Title4{}, err
	}
	st, err := state.NewTitle4()
	if err != nil {
		return Title4{}, err
	}
	dt, err := st.Initialize()
	if err != nil {
		return Title4{}, err
	}
	t := Title4{
		drawer: dr,
		state:  st,
		data:   dt,
	}
	return t, nil
}

func (t Title4) Update() (Scene, error) {
	if err := t.state.Update(t.data); err != nil {
		return nil, err
	}
	if t.state.Selector() == state.TitleBack {
		nt, err := NewTitle2()
		if err != nil {
			return nil, fmt.Errorf("faile to go to stage select: %w", err)
		}
		return nt, nil
	}
	return t, nil
}

func (t Title4) Draw(s *ebiten.Image) error {
	if err := t.drawer.Update(s, *t.data); err != nil {
		return err
	}
	return nil
}
