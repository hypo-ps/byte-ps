package ui

import "byte-ps/internal/terminal"

type Display struct {
	*terminal.RawTerminal
}

func NewDisplay(rawTerminal *terminal.RawTerminal) *Display {
	return &Display{
		RawTerminal: rawTerminal,
	}
}

func (d *Display) ClearScreen() {
	d.RawTerminal.ClearScreen()
}
