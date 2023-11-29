package main

import (
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

type Vector struct {
	X float64
	Y float64
}

type Game struct{
	playerPosition Vector
}

func (g *Game) Update() error {
	speed := float64(300 / ebiten.TPS())
	
	g.playerPosition.X += speed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// width := PlayerSprite.Bounds().Dx()
	// height := PlayerSprite.Bounds().Dy()
	// halfW := float64(width/2)
	// halfH := float64(height/2)

	op := &ebiten.DrawImageOptions{}
	// op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	// op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	// op.GeoM.Translate(halfW, halfH)

	// cm := colorm.ColorM{}
	// cm.Scale(1.0, 1.0, 1.0, 0.5)
	// colorm.DrawImage(screen, PlayerSprite, cm, op)
	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{playerPosition: Vector{X: 100, Y: 100}}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
