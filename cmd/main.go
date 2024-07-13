package main

import (
	"github.com/FluffyFoxTail/gorogue/game"
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gd := gamedata.NewGameData()
	g := game.NewGame(gd)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("AbobaRogue")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
