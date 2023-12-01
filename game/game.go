package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	meteors          []*Meteor
	bullets          []*Bullet
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(5 * time.Second),
	}

	g.player = NewPlayer(g)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for i, m := range g.meteors {
		m.Update()

		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
			}
		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	for _, b := range g.bullets {
		b.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
}
