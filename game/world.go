package game

import (
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type World struct {
	Manager    *ecs.Manager
	Player     *ecs.Component
	Renderable *ecs.Component
	Tags       map[string]*ecs.Tag
}

func InitializeWorld(startingLevel *level.Level) *World {
	tags := make(map[string]*ecs.Tag)
	manager := ecs.NewManager()

	playerAsset, _, err := ebitenutil.NewImageFromFile("assets/char.png")
	if err != nil {
		log.Fatal(err)
	}

	player := manager.NewComponent()
	renderable := manager.NewComponent()

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	playerEntity := &Player{Entity: &Entity{Image: playerAsset, X: x, Y: y}}

	manager.NewEntity().
		AddComponent(player, playerEntity).
		AddComponent(renderable, playerEntity)

	playersTag := ecs.BuildTag(player)
	tags["players"] = &playersTag

	renderables := ecs.BuildTag(renderable)
	tags["renderables"] = &renderables

	return &World{Manager: manager, Player: player, Renderable: renderable, Tags: tags}
}
