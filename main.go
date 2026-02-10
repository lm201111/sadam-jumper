package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
)

type Game struct{}

// Update — обновляет логику игры каждый кадр
func (g *Game) Update() error {
    return nil // пока логики нет
}

// Draw — отрисовывает каждый кадр
func (g *Game) Draw(screen *ebiten.Image) {
    ebitenutil.DebugPrint(screen, "Hello, Sadam Jumper!")
}

// Layout — задаёт размер окна
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 640, 480
}

func main() {
    game := &Game{}
    ebiten.SetWindowSize(640, 480)             // размер окна
    ebiten.SetWindowTitle("Sadam Jumper")      // название окна
    if err := ebiten.RunGame(game); err != nil { // запускаем игру
        log.Fatal(err)
    }
}
