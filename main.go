/*
ld46, a game made for Ludum Dare 46
Copyright (C) 2020  Loïg Jezequel

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
package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/loig/ld46/game"
	"github.com/loig/ld46/level"
)

var g game.Game

func init() {
	g.GameState = game.InLevel
	g.PlayerState = game.Alive
	g.ResetLevel = level.TestLevel
	g.CurrentLevel = g.ResetLevel.CopyLevel()
	img, _, err := ebitenutil.NewImageFromFile("images/tiles.png", ebiten.FilterDefault)
	g.PlayerX = g.CurrentLevel.PlayerInitialX
	g.PlayerY = g.CurrentLevel.PlayerInitialY
	g.FlowerState = g.CurrentLevel.FlowerInitialState
	if err != nil {
		panic(err)
	}
	g.Tiles = img
}

func main() {
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle(game.Title)
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
