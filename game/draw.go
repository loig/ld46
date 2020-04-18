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
)

//Draw for ensuring that Game implements the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	for y := 0; y < g.CurrentLevel.Height; y++ {
		for x := 0; x < g.CurrentLevel.Width; x++ {
			if g.CurrentLevel.Grid[y][x].IsFloor {
				tile := image.Rect(
					0, 0,
					16, 16,
				)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*16), float64(y*16))
				screen.DrawImage(g.Tiles.SubImage(tile).(*ebiten.Image), op)
			}
		}
	}

}
