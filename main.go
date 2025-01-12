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
	//game loop-while running , update state,draw screen ,else close
	running := false
	for {

		//draw logic
		// running logic

		event := screen.PollEvent()
	}

}
