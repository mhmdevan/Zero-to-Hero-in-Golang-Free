* * *

🐍 Snake Game in Go
===================

A simple implementation of the classic Snake game using Go and the Ebiten game library. The snake grows each time it eats food, and it wraps around the edges of the screen. The game ends if the snake collides with itself.

🎮 Features
-----------
s
*   **Grow the Snake**: Each time the snake eats food, its length increases.
*   **Screen Wrapping**: The snake reappears on the opposite side if it goes out of bounds.
*   **Score Tracking**: See your score as you grow the snake longer.
*   **Real-time Gameplay**: Smooth gameplay using the Ebiten game library.

🚀 Getting Started
------------------

### Prerequisites

1.  **Go**: Make sure you have Go installed. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2.  **Ebiten**: This project uses the Ebiten library for 2D graphics.

Install Ebiten by running:

bash

Copy code

`go get github.com/hajimehoshi/ebiten/v2`

### Installation

1.  Clone this repository:
    
2.  Initialize the Go module (if not already initialized):
    
    bash
    
    Copy code
    
    `go mod init snake-game go mod tidy`
    

### Running the Game

To run the game, execute:

bash

Copy code

`go run main.go snake.go food.go`

### Controls

*   **Arrow Keys**: Use the arrow keys to change the snake's direction.

### How to Play

*   Control the snake to eat food (red squares) that randomly appears on the screen.
*   Each time the snake eats food, its length increases.
*   Avoid colliding with yourself, as this will end the game.
*   The game wraps around the screen, so if you go off one edge, the snake will reappear on the opposite side.

🛠️ Project Structure
---------------------

bash

Copy code

`/snake-game │ ├── main.go             # Main game loop and input handling ├── snake.go            # Snake logic (movement, growth, collision detection) └── food.go             # Food logic (spawn and rendering)`

📄 Code Overview
----------------

### main.go

*   Initializes the game window and handles the main game loop.
*   Manages user input for controlling the snake.
*   Updates the game state (moving the snake, growing when food is eaten, checking for collisions).

### snake.go

*   Defines the `Snake` struct, which includes the snake's body, direction, and movement.
*   Handles snake growth, movement, and self-collision detection.
*   Implements screen wrapping, so the snake reappears on the opposite side if it exits the screen.

### food.go

*   Defines the `Food` struct, which randomly generates food on the screen.
*   Respawns food in a new location each time it's eaten.

📈 Game Mechanics
-----------------

*   **Growth**: Each time the snake eats food, a new segment is added to its tail.
*   **Collision Detection**: The game checks if the snake’s head overlaps any part of its body. If so, the game ends.
*   **Screen Wrapping**: If the snake goes beyond the screen bounds, it reappears on the opposite side.

📝 Future Improvements
----------------------

*   **Add High Scores**: Track the highest score across games.
*   **Add Obstacles**: Introduce static obstacles that the snake must avoid.
*   **Difficulty Levels**: Increase the speed of the snake as it grows or introduce different game modes.

📜 License
----------

This project is licensed under the MIT License.

* * *

Enjoy playing the Snake game and feel free to contribute with more features or improvements! 🐍

* * *
