package game

import (
	"time"
	
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth = 800
	ScreenHeight = 600
)

type Game struct{
	player *Player
	meteorSpawnTimer *Timer
	meteors []*Meteor
}

func NewGame() *Game {
	return &Game{player: NewPlayer(), meteorSpawnTimer: NewTimer(5 * time.Second)}
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	g := &Game{player: NewPlayer(), meteorSpawnTimer: NewTimer(5 * time.Second)}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}