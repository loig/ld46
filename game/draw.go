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
	"fmt"
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/loig/ld46/level"
)

//Draw for ensuring that Game implements the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	if g.GameState == BeginMenu {
		var button image.Rectangle
		// Play button
		button = image.Rect(
			0, 112,
			48, 128,
		)
		if g.MenuFocus == Play {
			button = image.Rect(
				48, 112,
				96, 128,
			)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(104, 128)
		screen.DrawImage(g.Tiles.SubImage(button).(*ebiten.Image), op)
		// Info button
		button = image.Rect(
			0, 128,
			48, 144,
		)
		if g.MenuFocus == Info {
			button = image.Rect(
				48, 128,
				96, 144,
			)
		}
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(104, 144)
		screen.DrawImage(g.Tiles.SubImage(button).(*ebiten.Image), op)
		// Quit button
		button = image.Rect(
			0, 144,
			48, 160,
		)
		if g.MenuFocus == Quit {
			button = image.Rect(
				48, 144,
				96, 160,
			)
		}
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(104, 160)
		screen.DrawImage(g.Tiles.SubImage(button).(*ebiten.Image), op)
		return
	}

	if g.GameState == InfoPage {
		text.Draw(screen, InfoBlock, g.DisplayFont, 30, 30, color.White)
		text.Draw(screen, InfoBlock2, g.DisplayFont, 30, 85, color.White)
		ccbysa := image.Rect(
			160, 224,
			256, 256,
		)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 90)
		screen.DrawImage(g.Tiles.SubImage(ccbysa).(*ebiten.Image), op)
		text.Draw(screen, InfoBlock3, g.DisplayFont, 10, 150, color.White)
	}

	if g.GameState == InLevel ||
		(g.GameState == LevelFinished &&
			(endLevelStepName(g.EndLevelStep) == congrats ||
				endLevelStepName(g.EndLevelStep) == fadeout)) {
		for y := 0; y < g.CurrentLevel.Height; y++ {
			for x := 0; x < g.CurrentLevel.Width; x++ {
				if g.CurrentLevel.FloorGrid[y][x].IsFloor {
					var tile image.Rectangle
					if g.CurrentLevel.FloorGrid[y][x].IsFallingTile {
						tileLife := g.CurrentLevel.FloorGrid[y][x].FallingTileLife
						tile = image.Rect(
							tileLife*16, 0,
							16+tileLife*16, 16,
						)
					} else if g.CurrentLevel.FloorGrid[y][x].IsLinkedTile {
						tile = image.Rect(
							64, 0,
							80, 16,
						)
					} else {
						tile = image.Rect(
							0, 0,
							16, 16,
						)
					}
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*16), float64(y*16))
					screen.DrawImage(g.Tiles.SubImage(tile).(*ebiten.Image), op)
				} else if g.CurrentLevel.FloorGrid[y][x].IsFallingTile || g.CurrentLevel.FloorGrid[y][x].IsLinkedTile {
					tilePos := level.TilePosition{TileX: x, TileY: y}
					animationStep, exists := g.FallingTilesAnimationStep[tilePos]
					if exists {
						var tile image.Rectangle
						if g.CurrentLevel.FloorGrid[y][x].IsFallingTile {
							tile = image.Rect(
								80+16*animationStep, 0,
								96+16*animationStep, 16,
							)
						} else { //LinkedTile
							tile = image.Rect(
								128+16*animationStep, 0,
								144+16*animationStep, 16,
							)
						}
						op := &ebiten.DrawImageOptions{}
						op.GeoM.Translate(float64(x*16), float64(y*16))
						screen.DrawImage(g.Tiles.SubImage(tile).(*ebiten.Image), op)
					}
				}
				if g.CurrentLevel.ObjectsGrid[y][x] != level.None {
					var object image.Rectangle
					switch g.CurrentLevel.ObjectsGrid[y][x] {
					case level.Player:
						switch g.PlayerState {
						case HoldingNothing:
							object = image.Rect(
								0+16*g.PlayerAnimationStep, 16,
								16+16*g.PlayerAnimationStep, 48,
							)
						case HoldingWater:
							object = image.Rect(
								32+16*g.PlayerAnimationStep, 16,
								48+16*g.PlayerAnimationStep, 48,
							)
						case Dead:
							if g.PlayerAnimationStep < deathSteps.animationSteps {
								object = image.Rect(
									64+16*g.PlayerAnimationStep, 16,
									80+16*g.PlayerAnimationStep, 48,
								)
							} else {
								object = image.Rect(
									0, 0,
									0, 0,
								)
							}
						}
					case level.FlowerPot:
						object = image.Rect(
							0+16*g.FlowerAnimationStep, 48,
							16+16*g.FlowerAnimationStep, 80,
						)
					case level.FlowerBud:
						object = image.Rect(
							32+16*g.FlowerAnimationStep, 48,
							48+16*g.FlowerAnimationStep, 80,
						)
					case level.FlowerBaby:
						object = image.Rect(
							64+16*g.FlowerAnimationStep, 48,
							80+16*g.FlowerAnimationStep, 80,
						)
					case level.FlowerGrown:
						object = image.Rect(
							96+16*g.FlowerAnimationStep, 48,
							112+16*g.FlowerAnimationStep, 80,
						)
					case level.Water:
						object = image.Rect(
							0, 80,
							16, 128,
						)
					}
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(x*16), float64(y*16)-16)
					screen.DrawImage(g.Tiles.SubImage(object).(*ebiten.Image), op)

				}
			}
		}
	}

	if g.GameState == LevelFinished {
		switch endLevelStepName(g.EndLevelStep) {
		case fadeout:
			fmt.Println("fadeout")
		case score:
			text.Draw(screen, congratulationText, g.DisplayFont, 30, 50, color.White)
			for i := 0; i < g.EndLevelAnimationStep-1; i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(scores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
		case endLevelNumberOfSteps:
			text.Draw(screen, congratulationText, g.DisplayFont, 30, 50, color.White)
			for i := 0; i < len(scores); i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(scores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
		}
		return
	}

}
