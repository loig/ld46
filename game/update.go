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
		if g.levelComplete() {
			g.GameState = LevelFinished
		}

		if g.PlayerState != Dead {

			newPlayerX := g.PlayerX
			newPlayerY := g.PlayerY
			oldPlayerX := g.PlayerX
			oldPlayerY := g.PlayerY
			hasMoved := false

			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				newPlayerY++
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX, newPlayerY+1) {
					newPlayerY++
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				newPlayerY--
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX, newPlayerY-1) {
					newPlayerY--
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				newPlayerX++
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX+1, newPlayerY) {
					newPlayerX++
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				newPlayerX--
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX-1, newPlayerY) {
					newPlayerX--
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if hasMoved {
				g.updateLevelGrid(oldPlayerX, oldPlayerY)
				g.updatePlayerState(oldPlayerX, oldPlayerY, newPlayerX, newPlayerY)
				return nil
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.resetGame()
		}
		return nil
	}

	return nil
}

func (g *Game) isValidPosition(newPlayerX, newPlayerY int) bool {
	if newPlayerX < 0 {
		return false
	}
	if newPlayerY < 0 {
		return false
	}
	if newPlayerX >= g.CurrentLevel.Width {
		return false
	}
	if newPlayerY >= g.CurrentLevel.Height {
		return false
	}
	objectOnNewPosition := g.CurrentLevel.ObjectsGrid[newPlayerY][newPlayerX]
	if g.PlayerState != HoldingNothing {
		return objectOnNewPosition == level.None
	}
	return objectOnNewPosition == level.None || objectOnNewPosition == level.Water
}

func (g *Game) updateLevelGrid(oldPlayerX, oldPlayerY int) {
	if g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFallingTile {
		g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].FallingTileLife--
		if g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].FallingTileLife <= 0 {
			g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFloor = false
			g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsFallingTile = false
		}
	}
}

func (g *Game) updatePlayerState(oldPlayerX, oldPlayerY, newPlayerX, newPlayerY int) {
	g.CurrentLevel.ObjectsGrid[oldPlayerY][oldPlayerX] = level.None
	g.PlayerX = newPlayerX
	g.PlayerY = newPlayerY
	tileBelowPlayer := g.CurrentLevel.FloorGrid[g.PlayerY][g.PlayerX]
	objectWithPlayer := g.CurrentLevel.ObjectsGrid[g.PlayerY][g.PlayerX]
	switch {
	case !tileBelowPlayer.IsFloor:
		g.PlayerState = Dead
	case objectWithPlayer == level.Water:
		g.PlayerState = HoldingWater
	case g.PlayerState == HoldingWater:
		nextToFlower, flowerX, flowerY := g.CurrentLevel.IsPositionNextToFlower(newPlayerX, newPlayerY)
		if nextToFlower {
			g.FlowerState = g.CurrentLevel.ObjectsGrid[flowerY][flowerX].Grow()
			g.CurrentLevel.ObjectsGrid[flowerY][flowerX] = g.FlowerState
			g.PlayerState = HoldingNothing
		}
	}
	g.CurrentLevel.ObjectsGrid[newPlayerY][newPlayerX] = level.Player
}

func (g *Game) resetGame() {
	g.CurrentLevel = g.ResetLevel.CopyLevel()
	g.PlayerX = g.ResetLevel.PlayerInitialX
	g.PlayerY = g.ResetLevel.PlayerInitialY
	g.PlayerState = HoldingNothing
	g.FlowerState = g.ResetLevel.FlowerInitialState
}

func (g *Game) levelComplete() bool {
	return g.FlowerState.IsFullyGrown()
}
