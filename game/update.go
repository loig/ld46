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

		oldPlayerX := g.PlayerX
		oldPlayerY := g.PlayerY

		if g.PlayerState != Dead {
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				if g.PlayerY < g.CurrentLevel.Height-1 {
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.None
					g.PlayerY++
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.Player
				}
				updateLevelGrid(g, oldPlayerX, oldPlayerY)
				checkPlayerState(g)
				return nil
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				if g.PlayerY > 0 {
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.None
					g.PlayerY--
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.Player
				}
				updateLevelGrid(g, oldPlayerX, oldPlayerY)
				checkPlayerState(g)
				return nil
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				if g.PlayerX < g.CurrentLevel.Width-1 {
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.None
					g.PlayerX++
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.Player
				}
				updateLevelGrid(g, oldPlayerX, oldPlayerY)
				checkPlayerState(g)
				return nil
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				if g.PlayerX > 0 {
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.None
					g.PlayerX--
					g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX] = level.Player
				}
				updateLevelGrid(g, oldPlayerX, oldPlayerY)
				checkPlayerState(g)
				return nil
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			resetGame(g)
		}
	}

	return nil
}

func updateLevelGrid(g *Game, oldPlayerX int, oldPlayerY int) {
	if g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFallingTile {
		g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].FallingTileLife--
		if g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].FallingTileLife <= 0 {
			g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFloor = false
			g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFallingTile = false
		}
	}
}

func checkPlayerState(g *Game) {
	tileBelowPlayer := g.CurrentLevel.FloorGrid[g.PlayerY][g.PlayerX]
	switch {
	case !tileBelowPlayer.IsFloor:
		g.PlayerState = Dead
	}
}

func resetGame(g *Game) {
	g.CurrentLevel = g.ResetLevel.CopyLevel()
	g.PlayerX = g.ResetLevel.PlayerX
	g.PlayerY = g.ResetLevel.PlayerY
	g.PlayerState = Alive
}
