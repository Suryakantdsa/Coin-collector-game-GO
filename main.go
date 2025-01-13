package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, msg string) {
	for index, char := range msg {
		screen.SetContent(x+index, y, char, nil, tcell.StyleDefault)
	}

}
func setupCoin(level int) []*Sprite {
	coins := make([]*Sprite, level+2)
	for index := range level + 2 {
		coins[index] = NewSprite('0', rand.Intn(20), rand.Intn(20))
	}
	return coins
}
func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err)
	}

	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	player := NewSprite('@', 10, 10)
	// coins := []*Sprite{
	// 	NewSprite('0', 10, 26),
	// 	NewSprite('0', 40, 10),
	// 	NewSprite('0', 15, 53),
	// }
	coins := setupCoin(1)

	score := 0
	level := 1

	//game loop-while running , update state,draw screen ,else close
	running := true
	for running {

		//draw logic
		screen.Clear()
		player.Draw(screen)
		for _, coin := range coins {
			coin.Draw(screen)
		}

		// ui

		drawString(
			screen,
			1, 1, fmt.Sprint("Score: ", score),
		)
		drawString(
			screen,
			1, 2, fmt.Sprint("level: ", level),
		)

		screen.Show()

		// running logic
		//getting the event
		event := screen.PollEvent()

		// checking the event type
		playerMoved := false
		switch event := event.(type) {
		case *tcell.EventKey:

			//checking the key
			switch event.Rune() {
			case 'q':
				running = false
			case '8':
				player.y -= 1
				playerMoved = true
			case '2':
				player.y += 1
				playerMoved = true

			case '4':
				player.x -= 1
				playerMoved = true

			case '6':
				player.x += 1
				playerMoved = true

			}
		}
		if playerMoved {
			coinCollectedIndex := -1
			for index, coin := range coins {
				if coin.x == player.x && coin.y == player.y {
					coinCollectedIndex = index
					score++
				}
			}

			if coinCollectedIndex > -1 {
				coins[coinCollectedIndex] = coins[len(coins)-1]
				coins = coins[0 : len(coins)-1]

				if len(coins) == 0 {
					level++
					coins = setupCoin(level)
				}
			}

			// 0 1 2 3 4
			// 0 1 4 3 4
		}
	}

}
