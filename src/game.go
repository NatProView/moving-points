package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gameStart() ([]point, int) {
	var size, quantity int
	fmt.Println("Provide size of the area (in range [10; 100])")
	fmt.Scan(&size)
	for size < 10 || size > 100 {
		fmt.Println("Wrong input, try again")
		fmt.Scan(&size)
	}
	fmt.Printf("Provide number of points (in range [2; %v])\n", size/4)
	fmt.Scan(&quantity)
	for quantity < 2 || quantity > size/4 {
		fmt.Println("Wrong input, try again")
		fmt.Scan(&quantity)

	}

	points := []point{}

	for i := 0; i < quantity; i++ {
		fmt.Println("Provide character for point ", i+1)
		var znak string
		fmt.Scan(&znak)
		for len(znak) != 1 {
			fmt.Println("Wrong input, try again")
			fmt.Scan(&znak)
		}
		points = append(
			points,
			point{
				x:    rand.Intn((size - 2) + 1), //TODO: POPRAWIC ZEBY ZAWSZE
				y:    rand.Intn((size - 2) + 1), //RESPILO SIE W RAMCE
				char: znak,
			},
		)
	}
	return points, size
}

func game() bool {
	rand.Seed(time.Now().UnixNano())
	points, size := gameStart()
	isOn := true

	for isOn {
		clear()

		for l := 0; l < len(points); l++ {
			points[l].x += move()
			points[l].y += move()
			check(&points[l].x, &points[l].y, size)
			for k := 0; k < len(points); k++ {
				if (points[l].x == points[k].x && points[l].y == points[k].y) && k != l {
					isOn = false
					fmt.Printf("%s and %s have hit each other!\n", points[l].char, points[k].char)

					break
				}
			}
		}
		if isOn {
			draw(size, points)
		}
		time.Sleep(250 * time.Millisecond)

	}
	return endScreen()
}

func endScreen() bool {
	fmt.Println("Would you like to play again? [y/n]")
	var yn string
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("Wrong input, try again")
		fmt.Scan(&yn)
	}
	if yn == "n" {
		fmt.Println("Thanks for playing!")
		return false

	}
	return true
}
