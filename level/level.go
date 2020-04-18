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
	Grid   [][]Tile
	Width  int
	Height int
}

//Tile defines the state of one tile of a level
type Tile struct {
	IsFloor bool
}

//TestLevel is a dummy level for testing
var TestLevel Level = Level{
	Width:  16,
	Height: 16,
	Grid: [][]Tile{
		[]Tile{Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
		[]Tile{Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{false}, Tile{true}, Tile{true}},
	},
}
