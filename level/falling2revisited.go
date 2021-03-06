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

//falling2revisited a level auto-generated from a text file
var falling2revisited Level = Level{
Width: 16 ,
Height: 16 ,
PlayerInitialX: 7 ,
PlayerInitialY: 8 ,
FlowerX: 9 ,
FlowerY: 6 ,
FlowerInitialState: FlowerPot ,
LinkedTiles: []TilePosition{
TilePosition{TileX: 7 , TileY: 7 },
TilePosition{TileX: 9 , TileY: 7 },
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
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 3},
Tile{IsFloor: true},
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
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 2},
Tile{IsFloor: true, IsLinkedTile: true},
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
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 2},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
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
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
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
Water,
None,
None,
FlowerPot,
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
Player,
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
Water,
None,
None,
Water,
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
