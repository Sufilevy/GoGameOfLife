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
	pixelScale  = 10
)

var board Board
var nextBoard Board

func setup() {
	p5.Canvas(boardWidth*pixelScale, boardHeight*pixelScale)
	p5.Background(color.Gray{220})
	p5.RandomSeed(uint64(time.Now().UnixNano()))
}

func draw() {
	drawBoard()
	time.Sleep(time.Millisecond * 50)
	updateBoard()
}

func setInitialBoard() {
	args := os.Args[1:]
	fileName := "Empty"
	if len(args) == 1 {
		fileName = args[0]
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
