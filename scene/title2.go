package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/drawer"
	"lapis2411/button-sample/state"
)

var _ Scene = (*Title2)(nil)

type Title2 struct {
	drawer drawer.Title
	state  *state.Title2
	data   *data.Title
}

func NewTitle2() (Title2, error) {
	dr, err := drawer.NewTitle()
	if err != nil {
		return Title2{}, err
	}
	st, err := state.NewTitle2()
	if err != nil {
		return Title2{}, err
	}
	dt, err := st.Initialize()
	if err != nil {
		return Title2{}, err
	}
	t := Title2{
		drawer: dr,
		state:  st,
		data:   dt,
	}
	return t, nil
}

func (t Title2) Update() (Scene, error) {
	if err := t.state.Update(t.data); err != nil {
		return nil, err
	}
	if t.state.Selector() == state.TitleBack {
		nt, err := NewTitle()
		if err != nil {
			return nil, fmt.Errorf("faile to go to stage select: %w", err)
		}
		return nt, nil
	} else if t.state.Selector() == state.TitleNext {
		nt, err := NewTitle3()
		if err != nil {
			return nil, fmt.Errorf("faile to go to stage select: %w", err)
		}
		return nt, nil
	}
	return t, nil
}

func (t Title2) Draw(s *ebiten.Image) error {
	if err := t.drawer.Update(s, *t.data); err != nil {
		return err
	}
	return nil
}
