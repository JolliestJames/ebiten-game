package assets

import (
	"fmt"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("player.png");
var MeteorSprites = mustLoadImages("meteors")

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
