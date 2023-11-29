package main

import (
	"time"
	"embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
)

const (
	ScreenWidth = 800
	ScreenHeight = 600
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

type Player struct {
	position Vector
	rotation float64
	sprite *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite: sprite,
		rotation: float64(0),
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}

	// speed := float64(300 / ebiten.TPS())

	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	g.player.position.Y += speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	g.player.position.Y -= speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyD) {
	// 	g.player.position.X += speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyA) {
	// 	g.player.position.X -= speed
	// }
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

type Game struct{
	player *Player
	moveTimer *Timer
}

func (g *Game) Update() error {
	g.player.Update()

	g.moveTimer.Update()

	if g.moveTimer.IsReady() {
		g.moveTimer.Reset()

		g.player.position.X += 25
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	g := &Game{player: NewPlayer(), moveTimer: NewTimer(5 * time.Second)}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
