package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	body      []Position
	direction Direction
}

type Position struct {
	x, y int
}

func NewSnake() *Snake {
	return &Snake{
		body:      []Position{{5, 5}}, // Initial position
		direction: Right,
	}
}

func (s *Snake) HeadX() int {
	return s.body[0].x
}

func (s *Snake) HeadY() int {
	return s.body[0].y
}

func (s *Snake) ChangeDirection(newDirection Direction) {
	// Prevent reversing direction
	if (s.direction == Up && newDirection == Down) ||
		(s.direction == Down && newDirection == Up) ||
		(s.direction == Left && newDirection == Right) ||
		(s.direction == Right && newDirection == Left) {
		return
	}
	s.direction = newDirection
}

func (s *Snake) Move(gridWidth, gridHeight int) {
	head := s.body[0]
	var newHead Position

	// Determine the new head position based on the current direction
	switch s.direction {
	case Up:
		newHead = Position{head.x, (head.y - 1 + gridHeight) % gridHeight}
	case Down:
		newHead = Position{head.x, (head.y + 1) % gridHeight}
	case Left:
		newHead = Position{(head.x - 1 + gridWidth) % gridWidth, head.y}
	case Right:
		newHead = Position{(head.x + 1) % gridWidth, head.y}
	}

	// Move the snake by adding the new head and removing the last segment
	s.body = append([]Position{newHead}, s.body[:len(s.body)-1]...)
}

func (s *Snake) Grow() {
	// Add a new segment at the tail's position to grow the snake
	tail := s.body[len(s.body)-1]
	s.body = append(s.body, tail)
}

func (s *Snake) CollidesWithSelf() bool {
	// Only check for self-collision if the snake has grown enough
	if len(s.body) < 4 {
		return false
	}

	head := s.body[0]
	for _, part := range s.body[1:] {
		if part == head {
			return true
		}
	}
	return false
}
func (s *Snake) Draw(screen *ebiten.Image) {
	for _, part := range s.body {
		ebitenutil.DrawRect(screen, float64(part.x*cellSize), float64(part.y*cellSize), cellSize, cellSize, color.RGBA{0, 255, 0, 255})
	}
}
