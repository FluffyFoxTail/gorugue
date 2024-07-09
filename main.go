package main

import (
	"github.com/FluffyFoxTail/gorogue/gamedata"
	"github.com/FluffyFoxTail/gorogue/gamemap"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game holds all data about entire game
type Game struct {
	*gamedata.GameData
	*gamemap.GameMap
}

// NewGame creates a new Game Object and initializes the data
func NewGame(gd *gamedata.GameData) *Game {
	return &Game{GameData: gd, GameMap: gamemap.NewGameMap(gd)}
}

// Update is called each tic.
func (g *Game) Update() error {
	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.GameMap.Dungeons[0].Levels[0]
	for x := 0; x < g.GameData.ScreenWidth; x++ {
		for y := 0; y < g.GameData.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y, g.GameData)]
			options := &ebiten.DrawImageOptions{}

			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.GameData.TileWidth * g.GameData.ScreenWidth, g.GameData.TileHeight * g.GameData.ScreenHeight
}
func main() {
	gd := gamedata.NewGameData()
	g := NewGame(gd)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("AbobaRogue")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
