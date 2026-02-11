package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lm201111/sadam-jumper/internal"
	"log"
	"os"
	"math/rand"
	"time"
)

var (
	audioContext *audio.Context
	bgmPlayer    *audio.Player
)

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

func main() {
	rand.Seed(time.Now().UnixNano())
	initAudio()

	planeImg, _, err := ebitenutil.NewImageFromFile("assets/plane.png")
	if err != nil {
		log.Fatal(err)
	}

	skyscraperImg, _, err := ebitenutil.NewImageFromFile("assets/skyscraper.png")
	if err != nil {
		log.Fatal(err)
	}

	bgImg, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	game := internal.NewGame(planeImg, skyscraperImg, bgImg)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sadam Jumper")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
