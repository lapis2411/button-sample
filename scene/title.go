package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/data"
	"lapis2411/button-sample/drawer"
	"lapis2411/button-sample/state"
)

var _ Scene = (*Title)(nil)

type Title struct {
	drawer drawer.Title
	state  *state.Title
	data   *data.Title
}

func NewTitle() (Title, error) {
	dr, err := drawer.NewTitle()
	if err != nil {
		return Title{}, err
	}
	st, err := state.NewTitle()
	if err != nil {
		return Title{}, err
	}
	dt, err := st.Initialize()
	if err != nil {
		return Title{}, err
	}
	t := Title{
		drawer: dr,
		state:  st,
		data:   dt,
	}
	return t, nil
}

func (t Title) Update() (Scene, error) {
	if err := t.state.Update(t.data); err != nil {
		return nil, err
	}
	if t.state.Selector() == state.TitleNext {
		nt, err := NewTitle2()
		if err != nil {
			return nil, fmt.Errorf("faile to go to stage select: %w", err)
		}
		return nt, nil
	}
	return t, nil
}

func (t Title) Draw(s *ebiten.Image) error {
	if err := t.drawer.Update(s, *t.data); err != nil {
		return err
	}
	return nil
}
