package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Println("hop on to byte-ps... :)")

	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		fmt.Println("not executed in terminal")
	}

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Print("error getting the terminal in raw mode, err: %w\r\n", err)
		return
	}

	defer term.Restore(fd, oldState)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		term.Restore(fd, oldState)
		os.Exit(1)
	}()

	buff := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buff)
		if err != nil {
			fmt.Println("error while reading the buffer, err: %w", err)
			continue
		}

		ch := buff[0]
		fmt.Printf("you entered byte: %q\r\n", ch)
		if ch == 'q' {
			fmt.Printf("exiting...\r\n")
			return
		}
	}
}
