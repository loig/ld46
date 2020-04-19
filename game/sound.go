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

//SoundName give names to sounds
type SoundName int

const numSound = 12

//These constants register the different sounds that can
//be played
const (
	PlayerFallSound SoundName = iota
	TileFallSound
	WaterFlowerSound
	TakeWaterSound
	PlayerMoveSound
	VictorySound
	ScoreDisplaySound
	PlayerDashSound
	LevelBeginSound
	ResetSound
	MenuMoveSound
	MenuConfirmSound
)

//PlaySound plays one sound effect
func (g *Game) PlaySound(soundName SoundName) {
	g.SoundPlayers[soundName].Rewind()
	g.SoundPlayers[soundName].Play()
}
