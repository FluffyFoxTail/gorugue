package gamemap

import (
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
)

// Dungeon is container for all levels at particular dungeon in the game world
type Dungeon struct {
	Name   string
	Levels []*level.Level
}
