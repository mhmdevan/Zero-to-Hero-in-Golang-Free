package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 520
	screenHeight = 440
	cellSize     = 10
	moveDelay    = 5 // Delay in frames for slowing down movement
)

type Game struct {
	snake     *Snake
	food      *Food
	score     int
	moveTimer int
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		snake: NewSnake(),
		food:  NewFood(),
	}
}

func (g *Game) Update() error {
	// Handle player input for changing direction
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.snake.ChangeDirection(Left)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.snake.ChangeDirection(Right)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.snake.ChangeDirection(Up)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.snake.ChangeDirection(Down)
	}

	// Control snake speed by delaying movement
	g.moveTimer++
	if g.moveTimer < moveDelay {
		return nil
	}
	g.moveTimer = 0

	// Move the snake
	g.snake.Move(screenWidth/cellSize, screenHeight/cellSize)

	// Check if snake has eaten the food
	if g.snake.HeadX() == g.food.x && g.snake.HeadY() == g.food.y {
		g.snake.Grow() // Increase snake length
		g.food.Respawn()
		g.score++
		fmt.Println("Snake ate food! Score:", g.score) // Debugging output
	}

	// Check if the snake collides with itself
	if g.snake.CollidesWithSelf() {
		fmt.Println("Game Over: Snake collided with itself") // Debugging output
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)
	g.food.Draw(screen)
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
