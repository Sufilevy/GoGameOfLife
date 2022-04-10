package main

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/go-p5/p5"
)

func index(x, y int) int {
	return y*boardWidth + x
}

func cords(i int) (int, int) {
	return i % boardWidth, i / boardWidth
}

type Board [boardWidth * boardHeight]bool

func (b *Board) getPixel(x, y int) bool {
	return b[(y*boardWidth)+x]
}

func (b *Board) setPixel(x, y int, value bool) {
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

func getNeighbourCount(x, y int) int {
	var neighbourCount int

	if x > 0 {
		if board[index(x-1, y)] {
			neighbourCount++
		}
		if y > 0 && board[index(x-1, y-1)] {
			neighbourCount++
		}
	}
	if x < boardWidth-1 {
		if board[index(x+1, y)] {
			neighbourCount++
		}
		if y < boardHeight-1 && board[index(x+1, y+1)] {
			neighbourCount++
		}
	}
	if y > 0 {
		if board[index(x, y-1)] {
			neighbourCount++
		}
		if x < boardWidth-1 && board[index(x+1, y-1)] {
			neighbourCount++
		}
	}
	if y < boardHeight-1 {
		if board[index(x, y+1)] {
			neighbourCount++
		}
		if x > 0 && board[index(x-1, y+1)] {
			neighbourCount++
		}
	}

	return neighbourCount
}

func shouldLive(x, y int) bool {
	shouldLive := false

	livingNeighbours := getNeighbourCount(x, y)
	isAlive := board[index(x, y)]

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

func updateBoardCocurrent(numRoutines int) {
	var wg sync.WaitGroup
	routineSize := len(board) / numRoutines

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go updateBoardPart(i*routineSize, i*routineSize+routineSize, &wg)
	}

	wg.Wait()
	board = nextBoard
	println("")
}

func updateBoardPart(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Done %d - %d\n", start, end)

	for i := start; i < end; i++ {
		x, y := cords(i)

		nextBoard[index(x, y)] = shouldLive(x, y)
	}
}
