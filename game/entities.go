package game

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*Entity
}

type Entity struct {
	Image *ebiten.Image
	X, Y  int
}

// TODO here use struct with concurent access???
func (e *Entity) Move(gd *gamedata.GameData, l *level.Level, x, y int) {
	index := l.GetIndexFromXY(e.X+x, e.Y+y, gd)
	if !l.Tiles[index].Blocked {
		e.X += x
		e.Y += y
		l.PlayerVisible.Compute(l, e.X, e.Y, 8)
	}
}

// TODO also here use struct with concurent access???
func (e *Entity) Render(gd *gamedata.GameData, l *level.Level, screen *ebiten.Image) {
	index := l.GetIndexFromXY(e.X, e.Y, gd)
	tile := l.Tiles[index]
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
	screen.DrawImage(e.Image, options)
}
