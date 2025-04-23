package ui

import "byte-ps/internal/terminal"

type Display struct {
	terminal *terminal.RawTerminal
}

func NewDisplay(rawTerminal *terminal.RawTerminal) *Display {
	return &Display{
		terminal: rawTerminal,
	}
}

func (d *Display) ClearScreen() {
	d.terminal.ClearScreen()
}
