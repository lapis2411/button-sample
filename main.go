package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"lapis2411/button-sample/scene"
)

type Game struct {
	scene scene.Scene
}

func (g *Game) Update() error {
	var err error
	g.scene, err = g.scene.Update()
	return err
}

func (g *Game) Draw(s *ebiten.Image) {
	if err := g.scene.Draw(s); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func main() {
	sc, err := scene.NewTitle()
	if err != nil {
		log.Fatal(err)
	}
	g := Game{
		scene: sc,
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
