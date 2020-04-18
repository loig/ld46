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

	g.UpdateAnimation()

	if g.GameState == BeginMenu {
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.MenuFocus = (g.MenuFocus + 1) % EndMenu
			return nil
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.MenuFocus = (g.MenuFocus - 1)
			if g.MenuFocus < 0 {
				g.MenuFocus = EndMenu - 1
			}
			return nil
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			switch g.MenuFocus {
			case Play:
				g.GameState = InLevel
				g.ResetAnimation()
			case Info:
				g.GameState = InfoPage
				g.ResetAnimation()
			case Quit:
				return ErrEndGame
			}
			return nil
		}
		return nil
	}

	if g.GameState == InfoPage {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.GameState = BeginMenu
			g.ResetAnimation()
		}
		return nil
	}

	if g.GameState == InLevel {
		if g.levelComplete() {
			g.GameState = LevelFinished
			g.ResetAnimation()
			g.Scores[resets] -= g.Scores[deaths]
			g.Scores[playerSteps] -= g.Scores[playerDashes]
		}

		if g.PlayerState != Dead {

			newPlayerX := g.PlayerX
			newPlayerY := g.PlayerY
			oldPlayerX := g.PlayerX
			oldPlayerY := g.PlayerY
			hasMoved := false
			hasDashed := false

			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				newPlayerY++
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX, newPlayerY+1) {
					newPlayerY++
					hasDashed = true
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				newPlayerY--
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX, newPlayerY-1) {
					newPlayerY--
					hasDashed = true
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				newPlayerX++
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX+1, newPlayerY) {
					newPlayerX++
					hasDashed = true
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				newPlayerX--
				if g.PlayerState == HoldingNothing && ebiten.IsKeyPressed(ebiten.KeySpace) && g.isValidPosition(newPlayerX-1, newPlayerY) {
					newPlayerX--
					hasDashed = true
					hasMoved = true
				} else if g.isValidPosition(newPlayerX, newPlayerY) {
					hasMoved = true
				}
			}

			if hasMoved {
				if hasDashed {
					g.Scores[playerDashes]++
				}
				g.Scores[playerSteps]++
				g.updateLevelGrid(oldPlayerX, oldPlayerY)
				g.updatePlayerState(oldPlayerX, oldPlayerY, newPlayerX, newPlayerY)
				return nil
			}
		} else if g.PlayerAnimationStep >= deathSteps.animationSteps {
			if g.PlayerAnimationFrame >= waitAfterDeath {
				g.resetGame()
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.resetGame()
		}
		return nil
	}

	if g.GameState == LevelFinished {

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.EndLevelStep >= endLevelNumberOfSteps {
				g.setNextLevel()
			}
			g.EndLevelStep++
			g.EndLevelAnimationStep = 0
			g.EndLevelAnimationFrame = 0
		}

		return nil
	}

	if g.GameState == GameFinished {

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if g.EndGameAnimationStep >= endGameSteps.animationSteps {
				return ErrEndGame
			}
			g.EndGameAnimationStep = endGameSteps.animationSteps
			g.EndGameAnimationFrame = 0
		}
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
			tilePos := level.TilePosition{TileX: oldPlayerX, TileY: oldPlayerY}
			g.FallingTilesAnimationStep[tilePos] = 0
			g.FallingTilesAnimationFrame[tilePos] = 0
			g.Scores[destroyed]++
		}
	} else if g.CurrentLevel.FloorGrid[oldPlayerY][oldPlayerX].IsLinkedTile {
		for _, coord := range g.CurrentLevel.LinkedTiles {
			g.CurrentLevel.FloorGrid[coord.TileY][coord.TileX].IsFloor = false
			g.FallingTilesAnimationStep[coord] = 0
			g.FallingTilesAnimationFrame[coord] = 0
		}
		g.Scores[destroyed] += len(g.CurrentLevel.LinkedTiles)
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
		g.ResetAnimation()
		g.Scores[deaths]++
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
	g.Scores[resets]++
	g.CurrentLevel = g.ResetLevel.CopyLevel()
	g.PlayerX = g.ResetLevel.PlayerInitialX
	g.PlayerY = g.ResetLevel.PlayerInitialY
	g.PlayerState = HoldingNothing
	g.FlowerState = g.ResetLevel.FlowerInitialState
	g.GameState = InLevel
	g.ResetAnimation()
}

func (g *Game) levelComplete() bool {
	return g.FlowerState.IsFullyGrown()
}

func (g *Game) setNextLevel() {
	for i := 0; i < len(g.Scores); i++ {
		g.TotalScores[i] += g.Scores[i]
		g.Scores[i] = 0
	}
	nextLevel, gameFinished, isTuto, levelNumber, tutoNumber := level.GetLevel()
	g.GameState = InLevel
	if isTuto {
		g.GameState = InTuto
	}
	if gameFinished {
		g.GameState = GameFinished
	}
	g.PlayerState = HoldingNothing
	if nextLevel != nil {
		g.ResetLevel = *nextLevel
		g.CurrentLevel = g.ResetLevel.CopyLevel()
	}
	g.LevelNumber = levelNumber
	g.TutoNumber = tutoNumber
	g.PlayerX = g.CurrentLevel.PlayerInitialX
	g.PlayerY = g.CurrentLevel.PlayerInitialY
	g.FlowerState = g.CurrentLevel.FlowerInitialState
	g.FallingTilesAnimationStep = make(map[level.TilePosition]int)
	g.FallingTilesAnimationFrame = make(map[level.TilePosition]int)
	g.ResetAnimation()
}
