package main

import (
	"time"
	"embed"
	"image"
	_ "image/png"
	// "math"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
)

//go:embed assets/*
var assets embed.FS

var PlayerSprite = mustLoadimage("assets/player.png");

func mustLoadimage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(image)
}

type Timer struct {
	currentTicks int
	targetTicks int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks: int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

type Vector struct {
	X float64
	Y float64
}

type Game struct{
	playerPosition Vector
	moveTimer *Timer
}

func (g *Game) Update() error {
	speed := float64(300 / ebiten.TPS())
	
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerPosition.X += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerPosition.X -= speed
	}

	g.moveTimer.Update()

	if g.moveTimer.IsReady() {
		g.moveTimer.Reset()

		g.playerPosition.X += 25
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)

	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{playerPosition: Vector{X: 100, Y: 100}, moveTimer: NewTimer(5 * time.Second)}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
