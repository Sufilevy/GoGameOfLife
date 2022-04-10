package main

import (
	"fmt"

	"github.com/go-p5/p5"
)

func getBoardName() string {
	fmt.Print("Enter the name of the board: ")

	var input string
	fmt.Scanln(&input)

	if input == "" {
		input = "Empty"
	}
	return input
}

func leftButtonPressed() bool {
	return p5.Event.Mouse.Buttons.Contain(p5.ButtonLeft)
}

func rightButtonPressed() bool {
	return p5.Event.Mouse.Buttons.Contain(p5.ButtonRight)
}

func getMouseCords() (float64, float64) {
	return p5.Event.Mouse.Position.X, p5.Event.Mouse.Position.Y
}

func handleInput() {
	if p5.Event.Mouse.Pressed {
		switch {
		case rightButtonPressed():
			running = !running
		case leftButtonPressed():
			if !running {
				x, y := getMouseCords()
				pixelValue := board.getPixel(int(x/pixelScale), int(y/pixelScale))
				board.setPixel(int(x/pixelScale), int(y/pixelScale), !pixelValue)
			}
		}
	}
}
