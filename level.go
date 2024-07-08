package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type TileType int8

const (
	FLOOR TileType = iota
	CHAR  TileType = iota
	WALL  TileType = iota
)

type MapTile struct {
	PixelX  int // upper left corner for render purposes
	PixelY  int
	Blocked bool // is it should block player
	Image   *ebiten.Image
}

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

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() *GameData {
	return &GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
}

func (gd *GameData) CreateTiles() []*MapTile {
	tiles := make([]*MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
				wall := NewMapTale(WALL, x*gd.TileWidth, y*gd.TileHeight, true)
				tiles = append(tiles, wall)
			} else {
				floor := NewMapTale(FLOOR, x*gd.TileWidth, y*gd.TileHeight, false)
				tiles = append(tiles, floor)
			}
		}
	}
	return tiles
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (gd *GameData) GetIndexFromXY(x int, y int) int {
	return (y * gd.ScreenWidth) + x
}
