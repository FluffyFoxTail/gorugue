package gamemap

import (
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
)

// GameMap holds all levels and information about current game session
type GameMap struct {
	Dungeons     []*Dungeon
	CurrentLevel *level.Level
}

// NewGameMap creates a new set of maps for the entire game.
func NewGameMap(l ...*level.Level) *GameMap {
	return &GameMap{
		Dungeons: []*Dungeon{
			{Name: "sample", Levels: l},
		},
		CurrentLevel: l[0],
	}
}
