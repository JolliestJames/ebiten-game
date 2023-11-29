package main

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
)

const (
	ScreenWidth = 800
	ScreenHeight = 600
)

//go:embed assets/*
var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png");
var MeteorSprites = mustLoadImages("assets/meteors")

func mustLoadImage(name string) *ebiten.Image {
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

func mustLoadImages(name string) []*ebiten.Image {
	d, err := assets.ReadDir(name)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(d))
	for i := range d {
		images[i] = mustLoadImage(fmt.Sprintf("%s/%s", name, d[i].Name()))
	}

	return images
}

type Meteor struct {
	position Vector
	rotation float64
	rotationSpeed float64
	movement Vector
	sprite *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	r := ScreenWidth / 2.0

	angle := rand.Float64() * 2 * math.Pi

	pos := Vector{
		X: target.X + math.Cos(angle) * r,
		Y: target.Y + math.Sin(angle) * r,
	}

	velocity := 0.25 + rand.Float64() * 1.5

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Meteor{
		position: pos,
		rotation: float64(0),
		rotationSpeed: -0.02 + rand.Float64()*0.04,
		movement: movement,
		sprite: sprite,
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
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

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return Vector{
		X: v.X / magnitude,
		Y: v.Y / magnitude,
	}
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
	meteorSpawnTimer *Timer
	meteors []*Meteor
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
