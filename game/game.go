package game

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamemap"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game holds all data about entire game
type Game struct {
	*gamedata.GameData
	*gamemap.GameMap
	World *World
	MoveState
	MoveCount int
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	gd := gamedata.NewGameData()
	l := level.NewLevel(gd)
	gmap := gamemap.NewGameMap(l)
	world := InitializeWorld(gmap.CurrentLevel)
	return &Game{GameData: gd, GameMap: gmap, World: world, MoveState: PlayerMove, MoveCount: 0}
}

// Update is called each tic.
func (g *Game) Update() error {
	g.MoveCount++
	if g.MoveState == PlayerMove && g.MoveCount > 10 {
		x, y := 0, 0
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			y = -1
			println("y-1")
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			y = 1
			println("y 1")
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			x = -1
			println("x-1")
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			x = 1
			println("x 1")
		}

		l := g.GameMap.CurrentLevel
		for _, result := range g.World.Manager.Query(*g.World.Tags["players"]) {
			player := result.Components[g.World.Player].(Movable)
			player.Move(g.GameData, l, x, y)
		}

		if x != 0 || y != 0 {
			g.MoveState = GetNextState(g.MoveState)
			g.MoveCount = 0
		}
	}
	g.MoveState = PlayerMove

	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	l := g.GameMap.CurrentLevel
	l.DrawLevel(g.GameData, screen)
	g.proccessRenderables(l, screen)
}

func (g *Game) proccessRenderables(l *level.Level, screen *ebiten.Image) {
	for _, result := range g.World.Manager.Query(*g.World.Tags["renderables"]) {
		entity := result.Components[g.World.Renderable].(Renderable)
		entity.Render(g.GameData, l, screen)
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.GameData.TileWidth * g.GameData.ScreenWidth, g.GameData.TileHeight * g.GameData.ScreenHeight
}
