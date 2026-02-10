package internal

import (
    "github.com/hajimehoshi/ebiten/v2"
    "math/rand"
)

type Obstacle struct {
    X, Y, Width, Height float64
}

const ObstacleSpeed = 3

func NewObstacle(x float64) *Obstacle {
    height := float64(100 + rand.Intn(200)) // случайная высота
    return &Obstacle{
        X:      x,
        Y:      480 - height,
        Width:  50,
        Height: height,
    }
}

func (o *Obstacle) Update() {
    o.X -= ObstacleSpeed
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
    rect := ebiten.NewImage(int(o.Width), int(o.Height))
    rect.Fill(color.RGBA{0x33, 0x99, 0xff, 0xff})
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(o.X, o.Y)
    screen.DrawImage(rect, opts)
}
