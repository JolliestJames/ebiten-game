package game

import (
	"math"
	"math/rand"

	"github.com/JolliestJames/ebiten-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	position      Vector
	rotation      float64
	rotationSpeed float64
	movement      Vector
	sprite        *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	r := ScreenWidth / 2.0

	angle := rand.Float64() * 2 * math.Pi

	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	velocity := 0.25 + rand.Float64()*1.5

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
		position:      pos,
		rotation:      float64(0),
		rotationSpeed: -0.02 + rand.Float64()*0.04,
		movement:      movement,
		sprite:        sprite,
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

func (m *Meteor) Collider() Rectangle {
	bounds := m.sprite.Bounds()

	return NewRectangle(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
