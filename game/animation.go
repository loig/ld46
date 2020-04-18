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

type animationInfo struct {
	framesPerAnimationStep int
	animationSteps         int
}

const (
	endLevelNumberOfSteps = 3
)

//End of level animation management
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
