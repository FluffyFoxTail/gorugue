package components

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamedata/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*Entity
}

type Entity struct {
	Image *ebiten.Image
	X, Y  int
}

func (e *Entity) Move(gd *gamedata.GameData, l *level.Level, x, y int) {
	index := l.GetIndexFromXY(x, y, gd)
	if !l.Tiles[index].Blocked {
		e.X = x
		e.Y = y
	}
}

func (e *Entity) Render(gd *gamedata.GameData, l *level.Level, screen *ebiten.Image) {
	index := l.GetIndexFromXY(e.X, e.Y, gd)
	tile := l.Tiles[index]
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
	screen.DrawImage(e.Image, options)
}
