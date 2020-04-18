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

//TestLevel is a dummy level for testing
var TestLevel Level = Level{
	Width:              16,
	Height:             16,
	PlayerInitialX:     5,
	PlayerInitialY:     4,
	FlowerX:            3,
	FlowerY:            3,
	FlowerInitialState: FlowerBud,
	LinkedTiles: []TilePosition{
		TilePosition{3, 6},
		TilePosition{5, 7},
		TilePosition{11, 5},
	},
	ObjectsGrid: [][]Object{
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, Water, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, FlowerBud, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, Player, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, Water, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, Water, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, Water, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
		[]Object{None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None},
	},
	FloorGrid: [][]Tile{
		[]Tile{Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, true, 2, false}, Tile{true, true, 3, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 2, false}, Tile{true, false, 0, false}, Tile{true, true, 3, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, true}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, true}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, true}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{false, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{false, false, 0, false}},
		[]Tile{Tile{true, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{true, false, 0, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}, Tile{true, true, 1, false}, Tile{true, false, 0, false}},
		[]Tile{Tile{true, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{false, false, 0, false}, Tile{true, false, 0, false}},
	},
}
