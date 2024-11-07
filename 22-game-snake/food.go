package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Food struct {
	x, y int
}

func NewFood() *Food {
	return &Food{x: rand.Intn(screenWidth / cellSize), y: rand.Intn(screenHeight / cellSize)}
}

func (f *Food) Respawn() {
	f.x = rand.Intn(screenWidth / cellSize)
	f.y = rand.Intn(screenHeight / cellSize)
}

func (f *Food) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(f.x*cellSize), float64(f.y*cellSize), cellSize, cellSize, color.RGBA{255, 0, 0, 255})
}
