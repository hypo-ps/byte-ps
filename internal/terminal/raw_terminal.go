package terminal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type RawTerminal struct {
	*Terminal
	isRaw    bool
	oldState *term.State
}

func NewRawTerminal() (*RawTerminal, error) {
	t, err := NewTerminal()
	if err != nil {
		return nil, err
	}

	return &RawTerminal{
		Terminal: t,
	}, nil
}

func (rt *RawTerminal) IsRaw() bool {
	return rt.isRaw
}

func (rt *RawTerminal) Begin() error {
	if rt.IsRaw() {
		return nil
	}

	oldState, err := term.MakeRaw(rt.fileDescriptor)
	if err != nil {
		return fmt.Errorf("error while starting raw mode, %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		term.Restore(rt.fileDescriptor, rt.oldState)
		os.Exit(1)
	}()

	rt.oldState = oldState
	rt.isRaw = true
	return nil
}

func (rt *RawTerminal) Close() error {
	if !rt.IsRaw() {
		return fmt.Errorf("terminal not in raw mode")
	}

	if err := term.Restore(rt.fileDescriptor, rt.oldState); err != nil {
		return fmt.Errorf("error while restoring the terminal, %v", err)
	}

	rt.oldState = nil
	rt.isRaw = false
	return nil
}
