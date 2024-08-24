package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() (Scene, error)
	Draw(s *ebiten.Image) error
}
