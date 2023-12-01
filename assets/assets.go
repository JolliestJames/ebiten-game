package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var BulletSprite = mustLoadImage("missile.png")
var PlayerSprite = mustLoadImage("player.png")
var MeteorSprites = mustLoadImages("meteors")
var ScoreFont = mustLoadFont("font.ttf")

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

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 48,
		DPI: 72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
} 