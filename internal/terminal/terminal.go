package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	fileDescriptor int
}

func NewTerminal() (*Terminal, error) {
	fd := int(os.Stdin.Fd())

	if !term.IsTerminal(fd) {
		return nil, fmt.Errorf("could not get terminal for fd %d", fd)
	}
	return &Terminal{
		fileDescriptor: fd,
	}, nil
}

func (t *Terminal) Read(b []byte) error {
	_, err := os.Stdin.Read(b)
	if err != nil {
		fmt.Printf("Error while reading the input, err: %v", err)
		return err
	}

	return nil
}

func (t *Terminal) Write(b []byte) {
	os.Stdout.Write(b)
}

func (t *Terminal) WriteLn(b []byte) {
	if b == nil {
		return
	}
	b = append(b, '\r', '\n')
	if _, err := os.Stdout.Write(b); err != nil {
		fmt.Printf("Error while writing, err: %v", err)
	}
}

func (t *Terminal) ClearScreen() {
	clearScreenStr := "\033[2J\033[H"
	t.Write([]byte(clearScreenStr))
}
