package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strconv"
)

type Game struct {
	Player        *Player
	PlaneImg      *ebiten.Image
	ExplosionImg  *ebiten.Image
	ObstacleImg   *ebiten.Image
	Background    *ebiten.Image
	Obstacles     []*Obstacle
	Score         int
	Tick          int
	GameOver      bool
	BgOffset      float64
}

func NewGame(plane, explosion, obstacle, bg *ebiten.Image) *Game {
	return &Game{
		Player:       NewPlayer(plane, 80, 200),
		PlaneImg:     plane,
		ExplosionImg: explosion,
		ObstacleImg:  obstacle,
		Background:   bg,
	}
}

func (g *Game) Update() error {

	if g.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame(g.PlaneImg, g.ExplosionImg, g.ObstacleImg, g.Background)
		}
		return nil
	}

	g.Tick++

	if ebiten.IsKeyPressed(ebiten.KeySpace) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.Player.Jump()
	}

	g.Player.Update()

	if g.Tick%120 == 0 {
		g.Obstacles = append(g.Obstacles, NewObstacle(g.ObstacleImg, 640))
	}

	for _, o := range g.Obstacles {
		o.Update()

		if g.Player.Rect().Overlaps(o.Rect()) {
			g.GameOver = true
		}

		if !o.Passed && o.X+float64(o.Width) < g.Player.X {
			g.Score++
			o.Passed = true
		}
	}

	if len(g.Obstacles) > 0 && g.Obstacles[0].X+float64(g.Obstacles[0].Width) < 0 {
		g.Obstacles = g.Obstacles[1:]
	}

	g.BgOffset -= 1
	if g.BgOffset <= -640 {
		g.BgOffset = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(g.BgOffset, 0)
	screen.DrawImage(g.Background, op1)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(g.BgOffset+1280, 0)
	screen.DrawImage(g.Background, op2)

	if g.GameOver {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.Player.X, g.Player.Y)
		screen.DrawImage(g.ExplosionImg, op)
	} else {
		g.Player.Draw(screen)
	}

	for _, o := range g.Obstacles {
		o.Draw(screen)
	}

	// Обычный счёт
	ebitenutil.DebugPrint(screen, "Счёт: "+strconv.Itoa(g.Score))

	if g.GameOver {
		ebitenutil.DebugPrintAt(screen, "ИГРА ОКОНЧЕНА", 220, 160)
		ebitenutil.DebugPrintAt(screen, "ФИНАЛЬНЫЙ СЧЁТ: "+strconv.Itoa(g.Score), 180, 200)
		ebitenutil.DebugPrintAt(screen, "Нажмите R для рестарта", 180, 240)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

