package level

import (
	"github.com/FluffyFoxTail/gorogue/gamedata"
	"github.com/FluffyFoxTail/gorogue/gamedata/level/maptile"
	"github.com/hajimehoshi/ebiten/v2"
)

// Level hold information about tile for dungeon level
type Level struct {
	Tiles []*maptile.MapTile
}

// NewLevel create a new game level in dungeon
func NewLevel(gd *gamedata.GameData) (level *Level) {
	level = &Level{}
	level.Tiles = level.CreateTiles(gd)
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
func (l *Level) CreateTiles(gd *gamedata.GameData) []*maptile.MapTile {
	tiles := make([]*maptile.MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
				wall := maptile.NewMapTale(maptile.WALL, x*gd.TileWidth, y*gd.TileHeight, true)
				tiles = append(tiles, wall)
			} else {
				floor := maptile.NewMapTale(maptile.FLOOR, x*gd.TileWidth, y*gd.TileHeight, false)
				tiles = append(tiles, floor)
			}
		}
	}
	return tiles
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (l *Level) GetIndexFromXY(x int, y int, gd *gamedata.GameData) int {
	return (y * gd.ScreenWidth) + x
}
