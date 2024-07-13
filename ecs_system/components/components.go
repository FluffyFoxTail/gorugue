package components

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamedata/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderable interface {
	Render(gd *gamedata.GameData, l *level.Level, screen *ebiten.Image)
}

type Movable interface {
	Move(gd *gamedata.GameData, l *level.Level, x, y int)
}
