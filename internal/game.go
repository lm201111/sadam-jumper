package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strconv"
)

type Game struct {
	Player      *Player
	Obstacles   []*Obstacle
	ObstacleImg *ebiten.Image
	Background  *ebiten.Image
	Score       int
	Tick        int
	GameOver    bool
	BgOffset    float64
}

func NewGame(playerImg, obstacleImg, bg *ebiten.Image) *Game {
	return &Game{
		Player:      NewPlayer(playerImg, 80, 200),
		ObstacleImg: obstacleImg,
		Background:  bg,
	}
}

func (g *Game) Update() error {

	if g.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame(g.Player.Img, g.ObstacleImg, g.Background)
		}
		return nil
	}

	g.Tick++

	// Прыжок
	if ebiten.IsKeyPressed(ebiten.KeySpace) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.Player.Jump()
	}

	g.Player.Update()

	// Спавн небоскрёбов
	if g.Tick%120 == 0 {
		g.Obstacles = append(g.Obstacles, NewObstacle(g.ObstacleImg, 640))
	}

	for _, o := range g.Obstacles {
		o.Update()

		// Столкновение
		if g.Player.Rect().Overlaps(o.Rect()) {
			g.GameOver = true
		}

		// Счёт
		if !o.Passed && o.X+float64(o.Width) < g.Player.X {
			g.Score++
			o.Passed = true
		}
	}

	// Удаление старых
	if len(g.Obstacles) > 0 && g.Obstacles[0].X+float64(g.Obstacles[0].Width) < 0 {
		g.Obstacles = g.Obstacles[1:]
	}

	// Прокрутка фона
	g.BgOffset -= 1
	if g.BgOffset <= -640 {
		g.BgOffset = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Рисуем фон два раза для бесконечной прокрутки
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(g.BgOffset, 0)
	screen.DrawImage(g.Background, op1)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(g.BgOffset+1280, 0)
	screen.DrawImage(g.Background, op2)

	g.Player.Draw(screen)

	for _, o := range g.Obstacles {
		o.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.Score))

	if g.GameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER", 260, 200)
		ebitenutil.DebugPrintAt(screen, "Press R to Restart", 220, 230)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
