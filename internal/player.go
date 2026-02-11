package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Player struct {
	X, Y      float64
	VelocityY float64
	Img       *ebiten.Image
	Width     int
	Height    int
}

const Gravity = 0.5
const JumpStrength = -10

func NewPlayer(img *ebiten.Image, startX, startY float64) *Player {
	w, h := img.Size()
	return &Player{
		X:      startX,
		Y:      startY,
		Img:    img,
		Width:  w,
		Height: h,
	}
}

func (p *Player) Update() {
	p.VelocityY += Gravity
	p.Y += p.VelocityY

	if p.Y > 480-float64(p.Height) {
		p.Y = 480 - float64(p.Height)
		p.VelocityY = 0
	}
}

func (p *Player) Jump() {
	p.VelocityY = JumpStrength
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Img, op)
}

func (p *Player) Rect() image.Rectangle {
	return image.Rect(int(p.X), int(p.Y), int(p.X)+p.Width, int(p.Y)+p.Height)
}
