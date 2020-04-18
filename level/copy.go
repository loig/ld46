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

// CopyLevel retruns a copy of the Level l
func (l Level) CopyLevel() Level {
	var copy Level

	copy.Width = l.Width
	copy.Height = l.Height
	copy.PlayerInitialX = l.PlayerInitialX
	copy.PlayerInitialY = l.PlayerInitialY
	copy.FlowerX = l.FlowerX
	copy.FlowerY = l.FlowerY
	copy.FlowerInitialState = l.FlowerInitialState

	copy.FloorGrid = make([][]Tile, l.Height)
	for y := 0; y < l.Height; y++ {
		copy.FloorGrid[y] = make([]Tile, l.Width)
		for x := 0; x < l.Width; x++ {
			copy.FloorGrid[y][x] = l.FloorGrid[y][x]
		}
	}

	copy.ObjectsGrid = make([][]Object, l.Height)
	for y := 0; y < l.Height; y++ {
		copy.ObjectsGrid[y] = make([]Object, l.Width)
		for x := 0; x < l.Width; x++ {
			copy.ObjectsGrid[y][x] = l.ObjectsGrid[y][x]
		}
	}

	copy.LinkedTiles = make([]TilePosition, len(l.LinkedTiles))
	for i := 0; i < len(l.LinkedTiles); i++ {
		copy.LinkedTiles[i] = l.LinkedTiles[i]
	}

	return copy
}
