package main

import (
	"fmt"
	"math/rand"
)

func check(x, y *int, size int) {
	switch *x {
	case 0:
		*x = size - 2
	case size - 1:
		*x = 1
	}
	switch *y {
	case 0:
		*y = size - 2
	case size - 1:
		*y = 1
	}
}

func move() int {
	switch rand.Intn(3) {
	case 0:
		return -1
	case 1:
		return 1
	case 2:
		return 0
	}
	fmt.Println("[where] something went wrong")
	return -10

}

func draw(size int, points []point) {
	tempBool := false
	tempChar := ""
	for i := 0; i < size; i++ { //y
		for j := 0; j < size; j++ { //x
			for _, n := range points {
				if n.x == j && n.y == i {
					tempBool = true
					tempChar = n.char
				}
			}
			if tempBool {
				fmt.Printf("%s", tempChar)
			} else if i == 0 || i == size-1 || j == 0 || j == size-1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
			tempBool = false
		}
		fmt.Printf("\n")
	}
}
