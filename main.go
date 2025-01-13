package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

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
	coins := []*Sprite{
		NewSprite('0', 10, 26),
		NewSprite('0', 40, 10),
		NewSprite('0', 15, 53),
	}
	//game loop-while running , update state,draw screen ,else close
	running := true
	for running {

		//draw logic
		screen.Clear()
		player.Draw(screen)
		for _, coin := range coins {
			coin.Draw(screen)
		}
		screen.Show()

		// running logic
		//getting the event
		event := screen.PollEvent()

		// checking the event type

		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Rune() {
			case 'q':
				running = false
			case '8':
				player.y -= 1
			case '2':
				player.y += 1
			case '4':
				player.x -= 1
			case '6':
				player.x += 1

			}
		}
	}

}
