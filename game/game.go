package game

import "github.com/hajimehoshi/ebiten"

type state int

//Constants defining the basic parameters of the game
const (
	Title        = "ld46"
	ScreenWidth  = 256
	ScreenHeight = 256
)

//Constants defining the possible states of the game
const (
	InLevel state = iota
)

//Game defines the general game structure
type Game struct {
	GameState state
}

//Update for ensuring that Game implements the ebiten.Game interface
func (g *Game) Update(screen *ebiten.Image) error {

	return nil
}

//Draw for ensuring that Game implements the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

}

//Layout for ensuring that Game implements the ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeigth int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
