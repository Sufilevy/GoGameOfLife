package main

import (
	"github.com/go-p5/p5"
	"image/color"
)

func index(x int, y int) int {
	return y*boardWidth + x
}

func cords(i int) (int, int) {
	return i % boardWidth, i / boardWidth
}

type Board [boardWidth * boardHeight]bool

func (b *Board) getPixel(x int, y int) bool {
	return b[(y*boardWidth)+x]
}

func (b *Board) setPixel(x int, y int, value bool) {
	b[(y*boardWidth)+x] = value
}

func getPixelValue(val bool) uint8 {
	if val {
		return 200
	}
	return 0
}

func drawBoard() {
	for i := 0; i < len(board); i++ {
		x, y := cords(i)

		pixelColor := color.Alpha{getPixelValue(board[i])}
		p5.Fill(pixelColor)
		p5.Stroke(pixelColor)
		p5.Square(float64(x*pixelScale), float64(y*pixelScale), pixelScale)
	}
}

func getNeighbourCount(x int, y int) int {
	var neighbourCount int

	if x > 0 {
		if board.getPixel(x-1, y) {
			neighbourCount++
		}
		if y > 0 && board.getPixel(x-1, y-1) {
			neighbourCount++
		}
	}
	if x < boardWidth-1 {
		if board.getPixel(x+1, y) {
			neighbourCount++
		}
		if y < boardHeight-1 && board.getPixel(x+1, y+1) {
			neighbourCount++
		}
	}
	if y > 0 {
		if board.getPixel(x, y-1) {
			neighbourCount++
		}
		if x < boardWidth-1 && board.getPixel(x+1, y-1) {
			neighbourCount++
		}
	}
	if y < boardHeight-1 {
		if board.getPixel(x, y+1) {
			neighbourCount++
		}
		if x > 0 && board.getPixel(x-1, y+1) {
			neighbourCount++
		}
	}

	return neighbourCount
}

func shouldLive(x int, y int) bool {
	shouldLive := false

	livingNeighbours := getNeighbourCount(x, y)
	isAlive := board.getPixel(x, y)

	if isAlive {
		switch {
		case livingNeighbours < 2:
			shouldLive = false
		case livingNeighbours > 3:
			shouldLive = false
		default:
			shouldLive = true
		}
	} else {
		if livingNeighbours == 3 {
			shouldLive = true
		} else {
			shouldLive = false
		}
	}

	return shouldLive
}

func updateBoard() {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			nextBoard.setPixel(x, y, shouldLive(x, y))
		}
	}
	board = nextBoard
}
