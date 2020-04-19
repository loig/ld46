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

//falling3 a level auto-generated from a text file
var falling3 Level = Level{
Width: 16 ,
Height: 16 ,
PlayerInitialX: 6 ,
PlayerInitialY: 7 ,
FlowerX: 8 ,
FlowerY: 8 ,
FlowerInitialState: FlowerBud ,
LinkedTiles: []TilePosition{
},
FloorGrid: [][]Tile{
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
},
ObjectsGrid: [][]Object{
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
Player,
None,
None,
None,
Water,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
Water,
None,
None,
FlowerBud,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
[]Object{
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
None,
},
},
}
