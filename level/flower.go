//Package level Copyright (C) 2020  Loïg Jezequel
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
package level

//Grow gives the next growing state of a plant
func (o Object) Grow() (grownFlower Object) {
	switch o {
	case FlowerPot:
		return FlowerBud
	case FlowerBud:
		return FlowerBaby
	case FlowerBaby:
		return FlowerGrown
	}
	return o
}

//IsFullyGrown tells if an object is a fully grown flower
func (o Object) IsFullyGrown() bool {
	return o == FlowerGrown
}

//IsPositionNextToFlower tells if the position x, y in the grid is right
//next to the flower or not
func (l Level) IsPositionNextToFlower(x, y int) (nextToFlower bool, flowerX, flowerY int) {
	return (x == l.FlowerX && y-1 == l.FlowerY) ||
		(x == l.FlowerX && y+1 == l.FlowerY) ||
		(x-1 == l.FlowerX && y == l.FlowerY) ||
		(x+1 == l.FlowerX && y == l.FlowerY), l.FlowerX, l.FlowerY
}
