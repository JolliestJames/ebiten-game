package game

import (
	"github.com/JolliestJames/ebiten-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Player struct {
	position      Vector
	rotation      float64
	sprite        *ebiten.Image
	shootCooldown Timer
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
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

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()
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
