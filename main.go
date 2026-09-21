package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"slices"

	player "simpleplayer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"secondgame/assets"
)

const (
	screenWidth  = 480
	screenHeight = 640
)

var lemy, sushiEgg, sushiKatsuo, sushiKazunoko, sushiSalmon *ebiten.Image

type BeltType string

type Transform struct {
	x      float64
	y      float64
	width  float64
	height float64
}

var Belts = struct {
	Up   BeltType
	Down BeltType
}{
	Up:   "Up",
	Down: "Down",
}

type Sushi struct {
	transform Transform
	image     *ebiten.Image
	id        int32
	belt      BeltType
}

type Obstacle struct {
	transform Transform
	image     *ebiten.Image
	id        int32
	belt      BeltType
}

type Game struct {
	player    player.Player
	sushi     []Sushi
	score     int32
	beltSpeed float64
	tick      int64
}

var loadedAssets *assets.Assets

func NewGame() *Game {
	g := &Game{}

	g.beltSpeed = 2.0

	g.sushi = make([]Sushi, 1, 10)
	g.player = player.Player{
		X:      (screenWidth - 32) / 2,
		Y:      (screenHeight - 40),
		Width:  32,
		Height: 32,
		Speed:  3.0,
	}

	g.sushi[0] = Sushi{
		transform: Transform{
			x:      (screenWidth - 32) / 2,
			y:      screenHeight * 0.3,
			width:  32,
			height: 32,
		},
		image: loadedAssets.Sushi.Egg,
		belt:  Belts.Up,
	}

	return g
}

func init() {

	player.InitSet(screenWidth, screenHeight)
	var err error

	loadedAssets, err = assets.Load()
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	g.tick++
	player.Move(&g.player)

	for i := range g.sushi {
		s := &g.sushi[i]

		//moving
		switch s.belt {
		case Belts.Up:
			s.transform.x += g.beltSpeed
			if s.transform.x >= screenWidth {
				s.belt = Belts.Down
				s.transform.y = screenHeight * 0.7
			}

		case Belts.Down:
			s.transform.x -= g.beltSpeed
			if s.transform.x+s.transform.width < 0 {
				s.belt = Belts.Up
				s.transform.y = screenHeight * 0.3
			}
		}

		//collision
		if player.CheckCollide(
			g.player.X, g.player.Y,
			g.player.Width, g.player.Height,
			s.transform.x, s.transform.y,
			s.transform.width, s.transform.height,
		) {
			//SYOUTOTU
			g.score += 100
			g.sushi = slices.Delete(g.sushi, i, i+1)
		}
	}

	if g.tick%10 == 0 {
		switch rand.Int32N(2) {
		case 0:
			//spawn Sushi
			sushi := &Sushi{}
			spawnOnBelt(&sushi.transform, &sushi.belt)
		case 1:
			//spawn Obstacle
		}
	}
	return nil
}

func spawnOnBelt(t *Transform, belt *BeltType) {
	switch rand.Int32N(2) {
	case 0:
		//Spawn on Upbelt
		// object.x = -object.width
		// object.y = screenHeight * 0.3
		// object.belt = Belts.Up
	case 1:
		//Spawn on Downbelt

	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// playerImg := ebiten.NewImage(int(g.player.Width), int(g.player.Height))
	// playerImg.Fill(color.RGBA{R: 200, A: 50})

	//lemy
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.2, 2.2)
	op.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(loadedAssets.Lemy, op)

	//sushi
	for _, s := range g.sushi {
		op.GeoM.Reset()
		op.GeoM.Scale(2.0, 2.0)
		op.GeoM.Translate(s.transform.x, s.transform.y)
		screen.DrawImage(s.image, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprint("Score: ", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sushi Party")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
