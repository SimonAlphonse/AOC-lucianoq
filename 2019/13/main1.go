package main

import "fmt"

type C struct{ x, y int }

type Tile int

// 0 is an empty tile. No game object appears in this tile.
// 1 is a wall tile. Walls are indestructible barriers.
// 2 is a block tile. Blocks can be broken by the ball.
// 3 is a horizontal paddle tile. The paddle is indestructible.
// 4 is a ball tile. The ball moves diagonally and bounces off objects.
const (
	Empty Tile = iota
	Wall
	Block
	Paddle
	Ball
)

func (t Tile) String() string {
	switch t {
	case Empty:
		return "  "
	case Wall:
		return "██"
	case Block:
		return "▒▒"
	case Paddle:
		return "▂▂"
	case Ball:
		return "⚽"
	}
	return ""
}

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	grid := make(map[C]Tile)

	var blocks int

	for {
		x, open := <-output

		if !open {
			fmt.Println(blocks)
			return
		}

		y := <-output
		tile := Tile(<-output)

		if tile == Block {
			blocks++
		}

		grid[C{x, y}] = tile
	}
}
