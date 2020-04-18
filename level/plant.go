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
func (o Object) Grow() (grownPlant Object) {
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

//IsPlant tells if an object is a plant
func (o Object) IsPlant() bool {
	return o == FlowerPot || o == FlowerBud || o == FlowerBaby || o == FlowerGrown
}
