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

//bigMoveLink a level auto-generated from a text file
var bigMoveLink Level = Level{
Width: 16 ,
Height: 16 ,
PlayerInitialX: 2 ,
PlayerInitialY: 13 ,
FlowerX: 11 ,
FlowerY: 4 ,
FlowerInitialState: FlowerBaby ,
LinkedTiles: []TilePosition{
TilePosition{TileX: 2 , TileY: 2 },
TilePosition{TileX: 4 , TileY: 2 },
TilePosition{TileX: 6 , TileY: 2 },
TilePosition{TileX: 8 , TileY: 2 },
TilePosition{TileX: 10 , TileY: 2 },
TilePosition{TileX: 12 , TileY: 2 },
TilePosition{TileX: 13 , TileY: 3 },
TilePosition{TileX: 2 , TileY: 4 },
TilePosition{TileX: 13 , TileY: 5 },
TilePosition{TileX: 2 , TileY: 6 },
TilePosition{TileX: 13 , TileY: 7 },
TilePosition{TileX: 2 , TileY: 8 },
TilePosition{TileX: 13 , TileY: 9 },
TilePosition{TileX: 2 , TileY: 10 },
TilePosition{TileX: 13 , TileY: 11 },
TilePosition{TileX: 2 , TileY: 12 },
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
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
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
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsLinkedTile: true},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
},
[]Tile{
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: false},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},
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
FlowerBaby,
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
Water,
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
Player,
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