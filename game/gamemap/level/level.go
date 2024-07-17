package level

import (
	"github.com/FluffyFoxTail/gorogue/game/gamedata"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/dice"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level/maptile"
	"github.com/FluffyFoxTail/gorogue/game/gamemap/level/room"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	MIN_ROOM_SIZE   = 6
	MAX_ROOM_SIZE   = 10
	MAX_ROOMS_COUNT = 30
)

// Level hold information about tile for dungeon level
type Level struct {
	Tiles []*maptile.MapTile
	Rooms []*room.Rectangle
}

// NewLevel create a new game level in dungeon
func NewLevel(gd *gamedata.GameData) (level *Level) {
	level = &Level{}
	level.Rooms = make([]*room.Rectangle, 0)
	level.GenerateLevelTiles(gd)
	return
}

func (l *Level) DrawLevel(gd *gamedata.GameData, screen *ebiten.Image) {
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := l.Tiles[l.GetIndexFromXY(x, y, gd)]
			options := &ebiten.DrawImageOptions{}

			options.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, options)
		}
	}
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (l *Level) GenerateLevelTiles(gd *gamedata.GameData) {
	tiles := l.InitTiles(gd)
	l.Tiles = tiles
	isContainsRoom := false

	for idx := 0; idx < MAX_ROOMS_COUNT; idx++ {
		w := dice.GetRandomBetween(MIN_ROOM_SIZE, MAX_ROOM_SIZE)
		h := dice.GetRandomBetween(MIN_ROOM_SIZE, MAX_ROOM_SIZE)
		x := dice.GetDiceRoll(gd.ScreenWidth - w - 1)
		y := dice.GetDiceRoll(gd.ScreenHeight - h - 1)

		newRoom := room.NewRectangle(x, y, w, h)
		isCanAddOnMap := true

		for _, otherRoom := range l.Rooms {
			if newRoom.IsIntersect(otherRoom) {
				isCanAddOnMap = false
				break
			}
		}

		if isCanAddOnMap {
			l.createRoom(newRoom, gd)
			if isContainsRoom {
				nexX, nexY := newRoom.Center()
				prevX, prevY := l.Rooms[len(l.Rooms)-1].Center()

				coinFlip := dice.GetDiceRoll(2)

				if coinFlip == 2 {
					l.createHorizontalTunnel(prevX, nexX, prevY, gd)
					l.createVerticalTunnel(prevY, nexY, prevX, gd)
				} else {
					l.createHorizontalTunnel(prevX, nexX, nexY, gd)
					l.createVerticalTunnel(prevY, nexY, prevX, gd)
				}
			}
			l.Rooms = append(l.Rooms, newRoom)
			isContainsRoom = true
		}
	}
}

func (l *Level) InitTiles(gd *gamedata.GameData) []*maptile.MapTile {
	tiles := make([]*maptile.MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = l.GetIndexFromXY(x, y, gd)
			wall := maptile.NewMapTale(maptile.WALL, x*gd.TileWidth, y*gd.TileHeight, true)
			tiles[index] = wall
		}
	}
	return tiles
}

func (l *Level) createRoom(room *room.Rectangle, gd *gamedata.GameData) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.GetIndexFromXY(x, y, gd)
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor32.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

func (l *Level) createHorizontalTunnel(x1, x2, y int, gd *gamedata.GameData) {
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := l.GetIndexFromXY(x, y, gd)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/hallway32.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

func (l *Level) createVerticalTunnel(y1, y2, x int, gd *gamedata.GameData) {
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := l.GetIndexFromXY(x, y, gd)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/hallway32.png")
			if err != nil {
				log.Fatal(err)
			}
			l.Tiles[index].Image = floor
		}
	}
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (l *Level) GetIndexFromXY(x int, y int, gd *gamedata.GameData) int {
	return (y * gd.ScreenWidth) + x
}
