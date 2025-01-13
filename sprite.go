package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	char rune
	x    int
	y    int
}

func NewSprite(char rune, x, y int) *Sprite {
	return &Sprite{
		char: char,
		x:    x,
		y:    y,
	}
}

func (s *Sprite) Draw(screen tcell.Screen) {
	screen.SetContent(
		s.x,
		s.y,
		s.char,
		nil,
		tcell.StyleDefault,
	)

}
