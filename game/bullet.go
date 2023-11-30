package game

import (
	_ "image/png"

	"github.com/JolliestJames/ebiten-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	sprite *ebiten.Image
	position Vector
	rotation float64
}

func NewBullet(p Vector, r float64) *Bullet {
	return &Bullet{
		position: p,
		rotation: r,
		sprite: assets.BulletSprite,
	}
}
