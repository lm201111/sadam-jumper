package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Obstacle struct {
	X      float64
	Y      float64
	Img    *ebiten.Image
	Width  int
	Height int
	Passed bool
}

const ObstacleSpeed = 4

func NewObstacle(img *ebiten.Image, x float64) *Obstacle {
	w, h := img.Size()
	return &Obstacle{
		X:      x,
		Y:      480 - float64(h),
		Img:    img,
		Width:  w,
		Height: h,
	}
}

func (o *Obstacle) Update() {
	o.X -= ObstacleSpeed
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.X, o.Y)
	screen.DrawImage(o.Img, op)
}

func (o *Obstacle) Rect() image.Rectangle {
	return image.Rect(int(o.X), int(o.Y), int(o.X)+o.Width, int(o.Y)+o.Height)
}
