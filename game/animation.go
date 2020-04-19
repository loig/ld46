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

import "github.com/loig/ld46/level"

//UpdateAnimation determines the step of the animation of each element
// of the game
func (g *Game) UpdateAnimation() {
	switch g.GameState {
	case BeginMenu:
		return
	case InTuto:
		return
	case InLevel:
		g.UpdatePlayerAnimation()
		g.UpdateFlowerAnimation()
		g.UpdateFallingTilesAnimation()
		return
	case LevelFinished:
		g.UpdateEndLevelAnimation()
		return
	case GameFinished:
		g.UpdateEndGameAnimation()
		return
	case InfoPage:
		return
	}
}

//ResetAnimation restarts all the animations
func (g *Game) ResetAnimation() {
	g.PlayerAnimationStep = 0
	g.PlayerAnimationFrame = 0
	g.FlowerAnimationStep = 0
	g.FlowerAnimationFrame = 0
	g.EndLevelStep = 0
	g.EndLevelAnimationStep = 0
	g.EndLevelAnimationFrame = 0
	g.FallingTilesAnimationStep = make(map[level.TilePosition]int)
	g.FallingTilesAnimationFrame = make(map[level.TilePosition]int)
	g.EndGameAnimationStep = 0
	g.EndGameAnimationFrame = 0
}

type animationInfo struct {
	framesPerAnimationStep int
	animationSteps         int
}

//Menus animation management
const (
	menuPlayChoiceText = "Play"
	menuInfoChoiceText = "Info"
	menuQuitChoiceText = "Quit"
	menuCommandInfo    = "Arrows to move. Enter to select."
	infoCommandInfo    = "Enter to quit."
	InfoBlock          = "A game developped in 48h\n    by Loig Jezequel\n   for Ludum Dare 46!\n(see  https://ldjam.com)"
	InfoBlock2         = "All images and sound are"
	InfoBlock3         = "  Code source under GPL-3.0\n         available at\n https://github.com/loig/ld46"
	scoreCommandInfo   = "Enter to continue."
	endGameCommandInfo = "Enter to quit."
)

//End of game animation management
var endGameSteps = animationInfo{20, 6}

const endCongratulationText = "Wow, you finished the game!"
const finalCongratulationText = "See you!"

//UpdateEndGameAnimation determines the step of the final animation
func (g *Game) UpdateEndGameAnimation() {
	if g.EndGameAnimationStep < endGameSteps.animationSteps {
		g.EndGameAnimationFrame++
		if g.EndGameAnimationFrame >= endGameSteps.framesPerAnimationStep {
			g.EndGameAnimationFrame = 0
			g.EndGameAnimationStep++
			g.PlaySound(ScoreDisplaySound)
		}
	}
}

//End of level animation management
const endLevelNumberOfSteps = 3

type endLevelStepName int

const congratulationText = "You finished level "

const (
	congrats endLevelStepName = iota
	fadeout
	score
)

var endLevelSteps = [endLevelNumberOfSteps]animationInfo{
	animationInfo{15, 3},
	animationInfo{60, 1},
	animationInfo{20, 6},
}

//UpdateEndLevelAnimation determines the step of the animation transition
//between levels
func (g *Game) UpdateEndLevelAnimation() {
	if g.EndLevelStep < endLevelNumberOfSteps {
		g.EndLevelAnimationFrame++
		if g.EndLevelAnimationFrame >= endLevelSteps[g.EndLevelStep].framesPerAnimationStep {
			g.EndLevelAnimationFrame = 0
			g.EndLevelAnimationStep++
			if g.EndLevelStep == 2 {
				g.PlaySound(ScoreDisplaySound)
			}
			if g.EndLevelAnimationStep >= endLevelSteps[g.EndLevelStep].animationSteps {
				g.EndLevelAnimationStep = 0
				g.EndLevelStep++
				if g.EndLevelStep == 1 {
					g.PlaySound(VictorySound)
				}
			}
		}
	}
}

//Player animation management
//Death animation management
var deathSteps = animationInfo{15, 3}

const waitAfterDeath = 60

//Stand animation management
var standSteps = animationInfo{10, 2}

//UpdatePlayerAnimation determines the step of the animation of the player
func (g *Game) UpdatePlayerAnimation() {
	switch g.PlayerState {
	case Dead:
		g.PlayerAnimationFrame++
		if g.PlayerAnimationStep < deathSteps.animationSteps {
			if g.PlayerAnimationFrame >= deathSteps.framesPerAnimationStep {
				g.PlayerAnimationFrame = 0
				g.PlayerAnimationStep++
			}
		}
	case HoldingWater, HoldingNothing:
		g.PlayerAnimationFrame++
		if g.PlayerAnimationFrame >= standSteps.framesPerAnimationStep {
			g.PlayerAnimationFrame = 0
			g.PlayerAnimationStep++
			if g.PlayerAnimationStep >= standSteps.animationSteps {
				g.PlayerAnimationStep = 0
			}
		}
	}
}

//Flower animation management
var flowerSteps = animationInfo{18, 2}

//UpdateFlowerAnimation determines the step of the animation of the player
func (g *Game) UpdateFlowerAnimation() {
	g.FlowerAnimationFrame++
	if g.FlowerAnimationFrame >= flowerSteps.framesPerAnimationStep {
		g.FlowerAnimationFrame = 0
		g.FlowerAnimationStep++
		if g.FlowerAnimationStep >= flowerSteps.animationSteps {
			g.FlowerAnimationStep = 0
		}
	}
}

//Falling tiles animation management
var fallingTileSteps = animationInfo{10, 3}

//UpdateFallingTilesAnimation determines the step of the animation of
//every currently falling tile
func (g *Game) UpdateFallingTilesAnimation() {
	for tilePos := range g.FallingTilesAnimationFrame {
		g.FallingTilesAnimationFrame[tilePos]++
		if g.FallingTilesAnimationFrame[tilePos] >= fallingTileSteps.framesPerAnimationStep {
			g.FallingTilesAnimationFrame[tilePos] = 0
			g.FallingTilesAnimationStep[tilePos]++
			if g.FallingTilesAnimationStep[tilePos] >= fallingTileSteps.animationSteps {
				delete(g.FallingTilesAnimationStep, tilePos)
				delete(g.FallingTilesAnimationFrame, tilePos)
				g.CurrentLevel.FloorGrid[tilePos.TileY][tilePos.TileX].IsFallingTile = false
				g.CurrentLevel.FloorGrid[tilePos.TileY][tilePos.TileX].IsLinkedTile = false
			}
		}
	}
}
