package editor

import (
	"byte-ps/internal/terminal"
	"byte-ps/internal/ui"
)

type Editor struct {
	terminal *terminal.RawTerminal
	display  *ui.Display
}

func NewEditor() (*Editor, error) {
	terminal, err := terminal.NewRawTerminal()
	if err != nil {
		return nil, err
	}
	return &Editor{
		terminal: terminal,
		display:  ui.NewDisplay(terminal),
	}, nil
}

func (e *Editor) Begin() error {
	err := e.terminal.Begin()
	if err != nil {
		return err
	}
	defer e.terminal.Close()

	e.display.ClearScreen()

	buff := make([]byte, 1)
	for {
		e.terminal.Read(buff)
		if buff[0] == 'q' {
			e.terminal.WriteLn([]byte("quiting..."))
			break
		}
		e.terminal.WriteLn(append([]byte("you enetered: "), buff...))
	}
	return nil
}
