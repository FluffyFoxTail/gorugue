package gamemap

import (
	"github.com/FluffyFoxTail/gorogue/gamedata"
	"github.com/FluffyFoxTail/gorogue/gamedata/dungeon"
	"github.com/FluffyFoxTail/gorogue/gamedata/level"
)

// GameMap holds all levels and information about current game session
type GameMap struct {
	Dungeons []*dungeon.Dungeon
}

// NewGameMap creates a new set of maps for the entire game.
// now for a single level
func NewGameMap(gd *gamedata.GameData) *GameMap {
	return &GameMap{
		Dungeons: []*dungeon.Dungeon{
			{Name: "sample", Levels: []*level.Level{level.NewLevel(gd)}},
		},
	}
}
