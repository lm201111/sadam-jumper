package internal

import (
    "github.com/hajimehoshi/ebiten/v2"
    "image"
)

type Player struct {
    X, Y      float64
    VelocityY float64
    Img       *ebiten.Image
}

const Gravity = 0.5
const JumpStrength = -10

func NewPlayer(img *ebiten.Image, startX, startY float64) *Player {
    return &Player{
        X:   startX,
        Y:   startY,
        Img: img,
    }
}

func (p *Player) Update() {
    // Гравитация
    p.VelocityY += Gravity
    p.Y += p.VelocityY

    // Пробегаем экран снизу
    if p.Y > 480-32 { // assuming height 480, plane 32px
        p.Y = 480 - 32
        p.VelocityY = 0
    }
}

func (p *Player) Jump() {
    p.VelocityY = JumpStrength
}

func (p *Player) Draw(screen *ebiten.Image) {
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(p.X, p.Y)
    screen.DrawImage(p.Img, opts)
}

func (p *Player) Rect() image.Rectangle {
    return image.Rect(int(p.X), int(p.Y), int(p.X)+32, int(p.Y)+32)
}
