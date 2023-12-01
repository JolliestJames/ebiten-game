package game

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rectangle) MaxX() float64 {
	return r.X + r.Width
}

func (r Rectangle) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rectangle) Intersects(other Rectangle) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}
