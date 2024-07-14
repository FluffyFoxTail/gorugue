package maptile

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// TileType enum for type of object on map
type TileType int8

const (
	FLOOR TileType = iota
	CHAR  TileType = iota
	WALL  TileType = iota
)

// MapTile represent each object on level
type MapTile struct {
	PixelX  int // upper left corner for render purposes
	PixelY  int
	Blocked bool // is it should block player
	Image   *ebiten.Image
}

// NewMapTale NewMapTile create new object for map of level objects
func NewMapTale(tileType TileType, pixelX, pixelY int, blocked bool) *MapTile {
	var file string
	switch tileType {
	case FLOOR:
		file = "assets/floor2.png"
	case CHAR:
		file = "assets/char.png"
	case WALL:
		file = "assets/wall2.png"
	}
	fromFile, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return &MapTile{
		PixelX:  pixelX,
		PixelY:  pixelY,
		Blocked: blocked,
		Image:   fromFile,
	}
}
