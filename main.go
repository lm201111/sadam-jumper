package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
    "math/rand"
    "time"
    "sadam-jumper/internal"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
    Player    *internal.Player
    Obstacles []*internal.Obstacle
    PlaneImg  *ebiten.Image
    Tick      int
}

func NewGame() *Game {
    img := ebiten.NewImage(32, 32)
    img.Fill(color.RGBA{0xff, 0, 0, 0xff}) // красный самолет
    return &Game{
        Player:   internal.NewPlayer(img, 50, 240),
        PlaneImg: img,
    }
}

func (g *Game) Update() error {
    g.Tick++

    // Прыжок: пробел или тап
    if ebiten.IsKeyPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        g.Player.Jump()
    }

    g.Player.Update()

    // Добавление новых небоскребов каждые 120 тиков (~2 сек)
    if g.Tick%120 == 0 {
        g.Obstacles = append(g.Obstacles, internal.NewObstacle(640))
    }

    // Обновляем и удаляем старые
    for i := 0; i < len(g.Obstacles); i++ {
        g.Obstacles[i].Update()
    }

    if len(g.Obstacles) > 0 && g.Obstacles[0].X+g.Obstacles[0].Width < 0 {
        g.Obstacles = g.Obstacles[1:]
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{0x88, 0xcc, 0xff, 0xff}) // фон

    g.Player.Draw(screen)
    for _, o := range g.Obstacles {
        o.Draw(screen)
    }

    ebitenutil.DebugPrint(screen, "Sadam Jumper")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 640, 480
}

func main() {
    rand.Seed(time.Now().UnixNano())
    game := NewGame()
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Sadam Jumper")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
