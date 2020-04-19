/*
Xtreme gardening 2020, a game created for Ludum Dare 46
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
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/loig/ld46/game"
	"golang.org/x/image/font"
)

var g game.Game

func init() {
	g.InitGame()

	img, _, err := ebitenutil.NewImageFromFile("images/tiles.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	g.Tiles = img

	ttfont, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		panic(err)
	}
	g.DisplayFont = truetype.NewFace(ttfont, &truetype.Options{
		Size:    8,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	g.InfoFont = truetype.NewFace(ttfont, &truetype.Options{
		Size:    6,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	musicFile, err := ebitenutil.OpenFile("sounds/music.mp3")
	music, err := mp3.Decode(g.AudioContext, musicFile)
	if err != nil {
		panic(err)
	}
	infiniteMusic := audio.NewInfiniteLoop(music, 70*4*44100)
	musicPlayer, err := audio.NewPlayer(g.AudioContext, infiniteMusic)
	if err != nil {
		panic(err)
	}
	g.MusicPlayer = musicPlayer

	soundFile, err := ebitenutil.OpenFile("sounds/playerFall.mp3")
	sound, err := mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err := audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.7)
	g.SoundPlayers[game.PlayerFallSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/tileFall.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.15)
	g.SoundPlayers[game.TileFallSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/waterFlower.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	g.SoundPlayers[game.WaterFlowerSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/takeWater.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	g.SoundPlayers[game.TakeWaterSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/playerMove.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.05)
	g.SoundPlayers[game.PlayerMoveSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/victory.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	g.SoundPlayers[game.VictorySound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/scoreDisplay.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.2)
	g.SoundPlayers[game.ScoreDisplaySound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/playerDash.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.15)
	g.SoundPlayers[game.PlayerDashSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/levelBegin.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	g.SoundPlayers[game.LevelBeginSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/reset.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	g.SoundPlayers[game.ResetSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/menuMove.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.05)
	g.SoundPlayers[game.MenuMoveSound] = player

	soundFile, err = ebitenutil.OpenFile("sounds/menuConfirm.mp3")
	sound, err = mp3.Decode(g.AudioContext, soundFile)
	if err != nil {
		panic(err)
	}
	player, err = audio.NewPlayer(g.AudioContext, sound)
	if err != nil {
		panic(err)
	}
	player.SetVolume(0.05)
	g.SoundPlayers[game.MenuConfirmSound] = player
}

func main() {
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle(game.Title)
	if err := ebiten.RunGame(&g); err != nil {
		if err != game.ErrEndGame {
			panic(err)
		}
	}
}
