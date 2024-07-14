package main

import (
	"github.com/FluffyFoxTail/gorogue/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("AbobaRogue")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
