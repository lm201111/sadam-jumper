package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/audio"
    "github.com/hajimehoshi/ebiten/v2/audio/vorbis"
    "github.com/lm201111/sadam-jumper/internal"
    "log"
    "math/rand"
    "os"
    "time"
    "image/color"
)

var (
    audioContext *audio.Context
    bgmPlayer    *audio.Player
)

// Инициализация фоновой музыки
func initAudio() {
    var err error
    audioContext = audio.NewContext(44100)
    f, err := os.Open("assets/music.ogg")
    if err != nil {
        log.Fatal(err)
    }
    d, err := vorbis.Decode(audioContext, f)
    if err != nil {
        log.Fatal(err)
    }
    bgmPlayer, err = audio.NewPlayer(audioContext, d)
    if err != nil {
        log.Fatal(err)
    }
    bgmPlayer.Play()
}

// Основная структура игры
type Game struct {
    Player    *internal.Player
    Obstacles []*internal.Obstacle
    PlaneImg  *ebiten.Image
    Tick      int
}

// Создание новой игры
func NewGame() *Game {
    img := ebiten.NewImage(32, 32)
    img.Fill(color.RGBA{0xff, 0, 0, 0xff}) // красный самолет
    return &Game{
        Player:   internal.NewPlayer(img, 50, 240),
        PlaneImg: img,
    }
}

// Логика игры
func (g *Game) Update() error {
    g.Tick++

    // Прыжок по пробелу или тапу
    if ebiten.IsKeyPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        g.Player.Jump()
    }

    g.Player.Update()

    // Добавляем новые небоскребы каждые 120 тиков (~2 сек)
    if g.Tick%120 == 0 {
        g.Obstacles = append(g.Obstacles, internal.NewObstacle(640))
    }

    // Обновляем и удаляем старые
    for _, o := range g.Obstacles {
        o.Update()
    }

    if len(g.Obstacles) > 0 && g.Obstacles[0].X+g.Obstacles[0].Width < 0 {
        g.Obstacles = g.Obstacles[1:]
    }

    return nil
}

// Отрисовка экрана
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{0x88, 0xcc, 0xff, 0xff}) // фон

    g.Player.Draw(screen)
    for _, o := range g.Obstacles {
        o.Draw(screen)
    }

    ebitenutil.DebugPrint(screen, "Sadam Jumper")
}

// Размер окна
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 640, 480
}

// main — запуск игры
func main() {
    rand.Seed(time.Now().UnixNano())
    initAudio() // запускаем музыку
    game := NewGame()
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Sadam Jumper")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
