package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Assets struct {
	Lemy  *ebiten.Image
	Sushi SushiImages
}

type SushiImages struct {
	Egg      *ebiten.Image
	Katsuo   *ebiten.Image
	Kazunoko *ebiten.Image
	Salmon   *ebiten.Image
	Tsuna    *ebiten.Image
}

func Load() (*Assets, error) {
	a := &Assets{}
	var err error

	a.Sushi.Egg, _, err = ebitenutil.NewImageFromFile("assets/image/sushi_egg.png")
	if err != nil {
		log.Fatal(err)
	}
	a.Sushi.Katsuo, _, err = ebitenutil.NewImageFromFile("assets/image/sushi_katsuo.png")
	if err != nil {
		log.Fatal(err)
	}
	a.Sushi.Kazunoko, _, err = ebitenutil.NewImageFromFile("assets/image/sushi_kazunoko.png")
	if err != nil {
		log.Fatal(err)
	}
	a.Sushi.Salmon, _, err = ebitenutil.NewImageFromFile("assets/image/sushi_salmon.png")
	if err != nil {
		log.Fatal(err)
	}
	a.Sushi.Tsuna, _, err = ebitenutil.NewImageFromFile("assets/image/sushi_tsuna.png")
	if err != nil {
		log.Fatal(err)
	}
	a.Lemy, _, err = ebitenutil.NewImageFromFile("assets/image/lemy.png")
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}
