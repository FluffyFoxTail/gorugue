package game

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
	"github.com/hajimehoshi/ebiten/v2"
)

// Renderable represent objects on game map that can be rendered
type Renderable interface {
	Render(gd *gamedata.GameData, l *level.Level, screen *ebiten.Image)
}

// Movable describe movable entities
type Movable interface {
	// Move try to reposition entity to Tile by x and y
	Move(gd *gamedata.GameData, l *level.Level, x, y int)
}
