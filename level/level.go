//Package level Copyright (C) 2020  Loïg Jezequel
/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package level

var currentLevel = 0
var currentTuto = 1

//Level defines the content of a single level of the game
type Level struct {
	FloorGrid          [][]Tile
	ObjectsGrid        [][]Object
	Width              int
	Height             int
	PlayerInitialX     int
	PlayerInitialY     int
	FlowerX            int
	FlowerY            int
	FlowerInitialState Object
	LinkedTiles        []TilePosition
}

//Tile defines the state of one tile of a level
type Tile struct {
	IsFloor         bool
	IsFallingTile   bool
	FallingTileLife int
	IsLinkedTile    bool
}

//TilePosition defines the position of a tile, used
//for linked tiles only
type TilePosition struct {
	TileX int
	TileY int
}

//Object defines an object of the game
type Object int

//Possible Objects
const (
	None Object = iota
	Player
	FlowerPot
	FlowerBud
	FlowerBaby
	FlowerGrown
	Water
)

//Order of levels
const (
	tuto1 int = iota
	afterMove
	tuto2
	afterFalling
	tuto3
	afterAged
	tuto4
	afterDash
	tuto5
	afterLinked
)

//GetLevel returns the next level to complete
func GetLevel() (level *Level, gameFinished, isTuto bool, levelNumber, tutoNumber int) {
	currentLevel++
	levelNumber = currentLevel - currentTuto + 1
	switch currentLevel - 1 {
	case tuto1, tuto2, tuto3, tuto4, tuto5:
		tutoNumber = currentTuto
		currentTuto++
		return nil, false, true, levelNumber, tutoNumber
	case afterMove:
		return &afterMoveTuto, false, false, levelNumber, currentTuto
	case afterDash:
		return &afterDashTuto, false, false, levelNumber, currentTuto
	case afterAged:
		return &afterAgedTuto, false, false, levelNumber, currentTuto
	case afterFalling:
		return &afterFallingTuto, false, false, levelNumber, currentTuto
	case afterLinked:
		return &afterLinkedTuto, false, false, levelNumber, currentTuto
	}
	return nil, true, false, 0, 0
}
