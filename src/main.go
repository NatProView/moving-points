package main

import (
	"os"
	"os/exec"
)

type point struct {
	x    int
	y    int
	char string
}

func main() {
	gameOn := true
	for gameOn {
		gameOn = game()

	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
