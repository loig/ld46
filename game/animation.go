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

//UpdateAnimation determines the step of the animation of each element
// of the game
func (g *Game) UpdateAnimation() {
	switch g.GameState {
	case BeginMenu:
	case InLevel:
		g.UpdatePlayerAnimation()
		g.UpdateFlowerAnimation()
	case LevelFinished:
		g.UpdateEndLevelAnimation()
	case GameFinished:
	case InfoPage:
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
}

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

//UpdateEndLevelAnimation determines the step of the animation transition
//between levels
func (g *Game) UpdateEndLevelAnimation() {
	if g.EndLevelStep < endLevelNumberOfSteps {
		g.EndLevelAnimationFrame++
		if g.EndLevelAnimationFrame >= endLevelSteps[g.EndLevelStep].framesPerAnimationStep {
			g.EndLevelAnimationFrame = 0
			g.EndLevelAnimationStep++
			if g.EndLevelAnimationStep >= endLevelSteps[g.EndLevelStep].animationSteps {
				g.EndLevelAnimationStep = 0
				g.EndLevelStep++
			}
		}
	}
}

type animationInfo struct {
	framesPerAnimationStep int
	animationSteps         int
}

//End of level animation management
const endLevelNumberOfSteps = 3

type endLevelStepName int

const congratulationText = "Congrats! You did it."

const (
	congrats endLevelStepName = iota
	fadeout
	score
)

var endLevelSteps = [endLevelNumberOfSteps]animationInfo{
	animationInfo{10, 5},
	animationInfo{10, 5},
	animationInfo{20, 6},
}

var scoreTexts = [5]string{
	"      Resets:   ",
	"      Deaths:   ",
	"Floor breaks:   ",
	"       Steps:   ",
	"      Dashes:   ",
}

type scoreName int

const (
	resets scoreName = iota
	deaths
	destroyed
	playerSteps
	playerDashes
)

var scores [5]int

//Death animation management
var deathSteps = animationInfo{15, 3}

const waitAfterDeath = 60

//Stand animation management
var standSteps = animationInfo{10, 2}

//Flower animation management
var flowerSteps = animationInfo{18, 2}
