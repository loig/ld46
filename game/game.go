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
	"github.com/loig/ld46/level"
)

//Constants defining the basic parameters of the game
const (
	Title        = "ld46"
	ScreenWidth  = 256
	ScreenHeight = 256
)

//State defines a possible state of the game
type State int

//Possible State values
const (
	InLevel State = iota
)

//PState defines a possible state of the player
type PState int

//Possible PState values
const (
	Dead PState = iota
	Alive
)

//Game defines the general game structure
type Game struct {
	GameState    State
	PlayerState  PState
	PlayerX      int
	PlayerY      int
	ResetLevel   level.Level
	CurrentLevel level.Level
	Tiles        *ebiten.Image
}

//Layout for ensuring that Game implements the ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
