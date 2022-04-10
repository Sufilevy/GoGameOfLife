package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/go-p5/p5"
)

const (
	boardWidth  = 80
	boardHeight = 60
	pixelScale  = 15
)

var board, nextBoard Board
var running = true

func setup() {
	p5.Canvas(boardWidth*pixelScale, boardHeight*pixelScale)
	p5.Background(color.Gray{220})
	p5.RandomSeed(uint64(time.Now().UnixNano()))
}

func draw() {
	handleInput()
	drawBoard()
	if running {
		updateBoard()
	}
}

func setInitialBoard() {
	args := os.Args[1:]
	var fileName string

	for _, arg := range args {
		switch arg[0] {
		case 'b':
			fileName = arg[2:]
		}
	}
	if fileName == "" {
		fileName = getBoardName()
	}

	file, err := os.Open(fmt.Sprintf("Boards/%s.txt", fileName))
	if err != nil {
		println("Board not found.")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, val := range scanner.Text() {
			if val == '1' {
				board.setPixel(x, y, true)
			}
		}
		y++
	}
}

func main() {
	setInitialBoard()
	p5.Run(setup, draw)
}
