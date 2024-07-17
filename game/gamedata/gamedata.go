package gamedata

// GameData hold size of elements on map
type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// NewGameData create fulfilled GameData struct
func NewGameData() *GameData {
	return &GameData{
		ScreenWidth:  40,
		ScreenHeight: 25,
		TileWidth:    32,
		TileHeight:   32,
	}
}
