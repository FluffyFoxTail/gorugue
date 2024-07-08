package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*GameData
	Tiles []*MapTile
}

func NewGame() *Game {
	gd := NewGameData()
	return &Game{GameData: gd, Tiles: gd.CreateTiles()}

}

// Update is called each tic.
func (g *Game) Update() error {
	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[g.GameData.GetIndexFromXY(x, y)]
			options := &ebiten.DrawImageOptions{}

			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 800
}
func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("AbobaRogue")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
