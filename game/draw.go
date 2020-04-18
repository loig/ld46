//Package game Copyright (C) 2020  Loïg Jezequel
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
package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/loig/ld46/level"
)

//Draw for ensuring that Game implements the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	if g.GameState == InLevel {
		for y := 0; y < g.CurrentLevel.Height; y++ {
			for x := 0; x < g.CurrentLevel.Width; x++ {
				if g.CurrentLevel.FloorGrid[y][x].IsFloor {
					var tile image.Rectangle
					if g.CurrentLevel.FloorGrid[y][x].IsFallingTile {
						tileLife := g.CurrentLevel.FloorGrid[y][x].FallingTileLife
						tile = image.Rect(
							tileLife*16, 0,
							16+tileLife*16, 16,
						)
					} else if g.CurrentLevel.FloorGrid[y][x].IsLinkedTile {
						tile = image.Rect(
							64, 0,
							80, 16,
						)
					} else {
						tile = image.Rect(
							0, 0,
							16, 16,
						)
					}
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*16), float64(y*16))
					screen.DrawImage(g.Tiles.SubImage(tile).(*ebiten.Image), op)
				}
				if g.CurrentLevel.ObjectsGrid[y][x] != level.None {
					var object image.Rectangle
					switch g.CurrentLevel.ObjectsGrid[y][x] {
					case level.Player:
						switch g.PlayerState {
						case HoldingNothing, Dead:
							object = image.Rect(
								0, 16,
								16, 48,
							)
						case HoldingWater:
							object = image.Rect(
								16, 16,
								32, 48,
							)
						}
					case level.FlowerPot:
						object = image.Rect(
							0, 48,
							16, 80,
						)
					case level.FlowerBud:
						object = image.Rect(
							16, 48,
							32, 80,
						)
					case level.FlowerBaby:
						object = image.Rect(
							32, 48,
							48, 80,
						)
					case level.FlowerGrown:
						object = image.Rect(
							48, 48,
							64, 80,
						)
					case level.Water:
						object = image.Rect(
							0, 80,
							16, 128,
						)
					}
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*16), float64(y*16)-16)
					screen.DrawImage(g.Tiles.SubImage(object).(*ebiten.Image), op)

				}
			}
		}
		return
	}

}
