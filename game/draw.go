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
					switch g.CurrentLevel.ObjectsGrid[y][x] {
					case level.Player:
						object := image.Rect(
							0, 16,
							16, 48,
						)
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*16), float64(y*16)-16)
						screen.DrawImage(g.Tiles.SubImage(object).(*ebiten.Image), op)
					}
				}
			}
		}
		return
	}

}
