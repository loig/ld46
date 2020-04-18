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
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/loig/ld46/level"
)

//Update for ensuring that Game implements the ebiten.Game interface
func (g *Game) Update(screen *ebiten.Image) error {

	if g.GameState == InLevel {
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			if g.CurrentLevel.PlayerY < g.CurrentLevel.Height-1 {
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.None
				g.CurrentLevel.PlayerY++
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.Player
			}
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			if g.CurrentLevel.PlayerY > 0 {
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.None
				g.CurrentLevel.PlayerY--
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.Player
			}
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			if g.CurrentLevel.PlayerX < g.CurrentLevel.Width-1 {
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.None
				g.CurrentLevel.PlayerX++
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.Player
			}
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if g.CurrentLevel.PlayerX > 0 {
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.None
				g.CurrentLevel.PlayerX--
				g.CurrentLevel.ObjectsGrid[g.CurrentLevel.PlayerY][g.CurrentLevel.PlayerX] = level.Player
			}
			return nil
		}
	}

	return nil
}
