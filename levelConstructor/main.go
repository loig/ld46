//Level constructor, generates levels from text
//Copyright (C) 2020  Loïg Jezequel
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
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	normalTile int = iota
	fallTile1
	fallTile2
	fallTile3
	linkTile
	noTile
)

const (
	nothing int = iota
	player
	water
	flowerPot
	flowerBud
	flowerBaby
	flowerGrown
)

const (
	fieldHeight = 16
	fieldWidth  = 16
)

var (
	playerX            int
	playerY            int
	flowerX            int
	flowerY            int
	flowerInitialState string
)

type lTilePos struct {
	x int
	y int
}

func main() {
	if len(os.Args) < 2 {
		panic("Please, enter file name as argument")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Please, entre correct file name as argument")
	}
	reader := bufio.NewReader(file)

	var floorGrid [][]int
	var objectsGrid [][]int
	var width int
	var height int
	var linkedTiles = make([]lTilePos, 0)

	for line, err := reader.ReadString('\n'); err != io.EOF; line, err = reader.ReadString('\n') {
		line = strings.Trim(line, "\n")
		if len(line) == 0 {
			break
		}
		if width == 0 {
			width = len(line)
		}
		height++
		if width != len(line) {
			s := fmt.Sprint("Not all lines have the same length, problem at line ", height)
			panic(s)
		}
		lineAsFloor := make([]int, len(line))
		for pos, char := range line {
			switch char {
			case '#':
				lineAsFloor[pos] = normalTile
			case '1':
				lineAsFloor[pos] = fallTile1
			case '2':
				lineAsFloor[pos] = fallTile2
			case '3':
				lineAsFloor[pos] = fallTile3
			case 'l', 'L':
				linkedTiles = append(linkedTiles, lTilePos{pos, height - 1})
				lineAsFloor[pos] = linkTile
			default:
				lineAsFloor[pos] = noTile
			}
		}
		floorGrid = append(floorGrid, lineAsFloor)
	}

	oHeight := 0

	heightAboveGrid := (fieldHeight - height) / 2
	widthLeftOfGrid := (fieldWidth - width) / 2

	for i := 0; i < len(linkedTiles); i++ {
		linkedTiles[i].x += widthLeftOfGrid
		linkedTiles[i].y += heightAboveGrid
	}

	for line, err := reader.ReadString('\n'); err != io.EOF; line, err = reader.ReadString('\n') {
		line = strings.Trim(line, "\n")
		oHeight++
		if width != len(line) {
			s := fmt.Sprint("Not all lines have the same length, problem in objects grid at line ", oHeight)
			panic(s)
		}
		lineAsObjects := make([]int, len(line))
		for pos, char := range line {
			switch char {
			case 'p', 'P':
				playerX = pos + widthLeftOfGrid
				playerY = oHeight - 1 + heightAboveGrid
				lineAsObjects[pos] = player
			case 'w', 'W':
				lineAsObjects[pos] = water
			case 'b':
				flowerX = pos + widthLeftOfGrid
				flowerY = oHeight - 1 + heightAboveGrid
				flowerInitialState = "FlowerPot"
				lineAsObjects[pos] = flowerPot
			case 'B':
				flowerX = pos + widthLeftOfGrid
				flowerY = oHeight - 1 + heightAboveGrid
				flowerInitialState = "FlowerBud"
				lineAsObjects[pos] = flowerBud
			case 'f':
				flowerX = pos + widthLeftOfGrid
				flowerY = oHeight - 1 + heightAboveGrid
				flowerInitialState = "FlowerBaby"
				lineAsObjects[pos] = flowerBaby
			case 'F':
				flowerX = pos + widthLeftOfGrid
				flowerY = oHeight - 1 + heightAboveGrid
				flowerInitialState = "FlowerGrown"
				lineAsObjects[pos] = flowerGrown
			default:
				lineAsObjects[pos] = nothing
			}
		}
		objectsGrid = append(objectsGrid, lineAsObjects)
	}

	if height != oHeight {
		panic("Floor grid and object grid do not have same size")
	}

	fmt.Println(
		"//Package level Copyright (C) 2020  Loïg Jezequel\n",
		"/*\n",
		"This program is free software: you can redistribute it and/or modify\n",
		"it under the terms of the GNU General Public License as published by\n",
		"the Free Software Foundation, either version 3 of the License, or\n",
		"(at your option) any later version.\n",
		"\n",
		"This program is distributed in the hope that it will be useful,\n",
		"but WITHOUT ANY WARRANTY; without even the implied warranty of\n",
		"MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the\n",
		"GNU General Public License for more details.\n",
		"\n",
		"You should have received a copy of the GNU General Public License\n",
		"along with this program.  If not, see <https://www.gnu.org/licenses/>.\n",
		"*/",
	)
	fmt.Println("package level")
	fmt.Println()
	fmt.Print("//", os.Args[1], " a level auto-generated from a text file\n")
	fmt.Println("var", os.Args[1], "Level = Level{")
	fmt.Println("Width:", fieldWidth, ",")
	fmt.Println("Height:", fieldHeight, ",")
	fmt.Println("PlayerInitialX:", playerX, ",")
	fmt.Println("PlayerInitialY:", playerY, ",")
	fmt.Println("FlowerX:", flowerX, ",")
	fmt.Println("FlowerY:", flowerY, ",")
	fmt.Println("FlowerInitialState:", flowerInitialState, ",")
	fmt.Println("LinkedTiles: []TilePosition{")
	for i := 0; i < len(linkedTiles); i++ {
		fmt.Println(
			"TilePosition{TileX:",
			linkedTiles[i].x, ", TileY:",
			linkedTiles[i].y, "},",
		)
	}
	fmt.Println("},")
	fmt.Println("FloorGrid: [][]Tile{")
	for i := 0; i < fieldHeight; i++ {
		fmt.Println("[]Tile{")
		for j := 0; j < fieldWidth; j++ {
			iInFloorGrid := i - heightAboveGrid
			jInFloorGrid := j - widthLeftOfGrid
			if i < heightAboveGrid || j < widthLeftOfGrid {
				fmt.Println("Tile{IsFloor: false},")
			} else if iInFloorGrid < height && jInFloorGrid < width {
				switch floorGrid[iInFloorGrid][jInFloorGrid] {
				case normalTile:
					fmt.Println("Tile{IsFloor: true},")
				case fallTile1:
					fmt.Println("Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 1},")
				case fallTile2:
					fmt.Println("Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 2},")
				case fallTile3:
					fmt.Println("Tile{IsFloor: true, IsFallingTile: true, FallingTileLife: 3},")
				case linkTile:
					fmt.Println("Tile{IsFloor: true, IsLinkedTile: true},")
				case noTile:
					fmt.Println("Tile{IsFloor: false},")
				}
			} else {
				fmt.Println("Tile{IsFloor: false},")
			}
		}
		fmt.Println("},")
	}
	fmt.Println("},")
	fmt.Println("ObjectsGrid: [][]Object{")
	for i := 0; i < fieldHeight; i++ {
		fmt.Println("[]Object{")
		for j := 0; j < fieldWidth; j++ {
			iInObjectsGrid := i - heightAboveGrid
			jInObjectsGrid := j - widthLeftOfGrid
			if i < heightAboveGrid || j < widthLeftOfGrid {
				fmt.Println("None,")
			} else if iInObjectsGrid < height && jInObjectsGrid < width {
				switch objectsGrid[iInObjectsGrid][jInObjectsGrid] {
				case nothing:
					fmt.Println("None,")
				case player:
					fmt.Println("Player,")
				case water:
					fmt.Println("Water,")
				case flowerPot:
					fmt.Println("FlowerPot,")
				case flowerBud:
					fmt.Println("FlowerBud,")
				case flowerBaby:
					fmt.Println("FlowerBaby,")
				case flowerGrown:
					fmt.Println("FlowerGrown,")
				}
			} else {
				fmt.Println("None,")
			}
		}
		fmt.Println("},")
	}
	fmt.Println("},")
	fmt.Println("}")
}
