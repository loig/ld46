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
		text.Draw(screen, menuPlayChoiceText, g.DisplayFont, 112, 139, color.White)
		text.Draw(screen, menuInfoChoiceText, g.DisplayFont, 112, 155, color.White)
		text.Draw(screen, menuQuitChoiceText, g.DisplayFont, 112, 171, color.White)
		text.Draw(screen, menuCommandInfo, g.InfoFont, 35, 250, color.White)
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
		text.Draw(screen, infoCommandInfo, g.InfoFont, 92, 250, color.White)
	}

	if g.GameState == InTuto {
		switch g.TutoNumber {
		case 1:
			text.Draw(screen, tuto1text1, g.DisplayFont, 30, 40, color.White)
			text.Draw(screen, tuto1text2, g.DisplayFont, 30, 90, color.White)
			text.Draw(screen, tuto1text3, g.DisplayFont, 30, 140, color.White)
			text.Draw(screen, tuto1text4, g.DisplayFont, 30, 190, color.White)
			toDraw := image.Rect(
				0, 16,
				16, 48,
			)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 20)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
			toDraw = image.Rect(
				0, 80,
				16, 112,
			)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(120, 65)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
			toDraw = image.Rect(
				64, 48,
				80, 80,
			)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 115)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
		case 2:
			text.Draw(screen, tuto2text1, g.DisplayFont, 30, 40, color.White)
			text.Draw(screen, tuto2text2, g.DisplayFont, 30, 90, color.White)
			text.Draw(screen, tuto2text3, g.DisplayFont, 30, 140, color.White)
			text.Draw(screen, tuto2text4, g.DisplayFont, 30, 190, color.White)
			toDraw := image.Rect(
				16, 0,
				32, 16,
			)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 82)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
		case 3:
			text.Draw(screen, tuto3text1, g.DisplayFont, 30, 80, color.White)
			text.Draw(screen, tuto3text2, g.DisplayFont, 30, 130, color.White)
		case 4:
			text.Draw(screen, tuto4text1, g.DisplayFont, 30, 40, color.White)
			text.Draw(screen, tuto4text2, g.DisplayFont, 30, 90, color.White)
			text.Draw(screen, tuto4text3, g.DisplayFont, 30, 140, color.White)
			text.Draw(screen, tuto4text4, g.DisplayFont, 30, 190, color.White)
			toDraw := image.Rect(
				48, 0,
				64, 16,
			)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 82)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
			toDraw = image.Rect(
				32, 0,
				48, 16,
			)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 132)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
			toDraw = image.Rect(
				16, 0,
				32, 16,
			)
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 182)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
		case 5:
			text.Draw(screen, tuto5text1, g.DisplayFont, 30, 40, color.White)
			text.Draw(screen, tuto5text2, g.DisplayFont, 30, 90, color.White)
			text.Draw(screen, tuto5text3, g.DisplayFont, 30, 140, color.White)
			toDraw := image.Rect(
				64, 0,
				80, 16,
			)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(190, 82)
			screen.DrawImage(g.Tiles.SubImage(toDraw).(*ebiten.Image), op)
		}
		text.Draw(screen, tutoCommandInfo, g.InfoFont, 75, 250, color.White)
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
						if g.GameState == LevelFinished {
							if endLevelStepName(g.EndLevelStep) == congrats {
								object = image.Rect(
									112+16*g.EndLevelAnimationStep, 16,
									128+16*g.EndLevelAnimationStep, 48,
								)
							} else {
								object = image.Rect(
									144, 16,
									160, 48,
								)
							}
						} else {
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
		case score:
			congratulationForDisplay := congratulationText + strconv.Itoa(g.LevelNumber)
			text.Draw(screen, congratulationForDisplay, g.DisplayFont, 30, 50, color.White)
			for i := 0; i < g.EndLevelAnimationStep-1; i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(g.Scores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
		case endLevelNumberOfSteps:
			congratulationForDisplay := congratulationText + strconv.Itoa(g.LevelNumber)
			text.Draw(screen, congratulationForDisplay, g.DisplayFont, 30, 50, color.White)
			for i := 0; i < len(g.Scores); i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(g.Scores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
			text.Draw(screen, scoreCommandInfo, g.InfoFont, 75, 250, color.White)
		}
		return
	}

	if g.GameState == GameFinished {
		text.Draw(screen, endCongratulationText, g.DisplayFont, 20, 50, color.White)
		if g.EndGameAnimationStep < endGameSteps.animationSteps {
			for i := 0; i < g.EndGameAnimationStep-1; i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(g.TotalScores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
		} else {
			for i := 0; i < len(g.TotalScores); i++ {
				scoreForDisplay := scoreTexts[i] + strconv.Itoa(g.TotalScores[i])
				text.Draw(screen, scoreForDisplay, g.DisplayFont, 40, 70+i*10, color.White)
			}
			text.Draw(screen, finalCongratulationText, g.DisplayFont, 150, 220, color.White)
			text.Draw(screen, endGameCommandInfo, g.InfoFont, 92, 250, color.White)
		}
	}

}
