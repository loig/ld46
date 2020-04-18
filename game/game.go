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
	"errors"

	"github.com/hajimehoshi/ebiten"
	"github.com/loig/ld46/level"
	"golang.org/x/image/font"
)

//Constants defining the basic parameters of the game
const (
	Title        = "ld46"
	ScreenWidth  = 256
	ScreenHeight = 256
	InfoBlock    = "A game developped in 48h\n    by Loig Jezequel\n   for Ludum Dare 46!\n(see  https://ldjam.com)"
	InfoBlock2   = "All images and sound are"
	InfoBlock3   = "  Code source under GPL-3.0\n         available at\n https://github.com/loig/ld46"
)

//ErrEndGame is raised at the end of a game
var ErrEndGame = errors.New("End")

//State defines a possible state of the game
type State int

//Possible State values
const (
	BeginMenu State = iota
	InTuto
	InLevel
	LevelFinished
	GameFinished
	InfoPage
)

//PState defines a possible state of the player
type PState int

//Possible PState values
const (
	Dead PState = iota
	HoldingNothing
	HoldingWater
)

//Focus defines the focused item of the menu
type Focus int

//Possible Focus values
const (
	Play Focus = iota
	Info
	Quit
	EndMenu
)

//Game defines the general game structure
type Game struct {
	GameState                  State
	PlayerAnimationStep        int
	PlayerAnimationFrame       int
	FlowerAnimationStep        int
	FlowerAnimationFrame       int
	EndLevelStep               int
	EndLevelAnimationStep      int
	EndLevelAnimationFrame     int
	FallingTilesAnimationStep  map[level.TilePosition]int
	FallingTilesAnimationFrame map[level.TilePosition]int
	PlayerState                PState
	PlayerX                    int
	PlayerY                    int
	FlowerState                level.Object
	ResetLevel                 level.Level
	CurrentLevel               level.Level
	LevelNumber                int
	TutoNumber                 int
	Tiles                      *ebiten.Image
	MenuFocus                  Focus
	DisplayFont                font.Face
}

//Layout for ensuring that Game implements the ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

//InitGame performs all necessary stuff before starting a game,
//except for media loading
func (g *Game) InitGame() {
	nextLevel, gameFinished, isTuto, levelNumber, tutoNumber := level.GetLevel()
	g.GameState = BeginMenu
	if isTuto {
		g.GameState = InTuto
	}
	if gameFinished {
		g.GameState = GameFinished
	}
	g.MenuFocus = Play
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
}
