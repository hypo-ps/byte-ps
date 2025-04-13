package main

import (
	"byte-ps/internal/terminal"
	"byte-ps/internal/ui"
	"fmt"
)

func main() {
	fmt.Println("hop on to byte-ps... :)")
	rawTerm, err := terminal.NewRawTerminal()
	if err != nil {
		fmt.Printf("failed to start the editor, %v\n", err)
	}
	defer rawTerm.Close()

	display := ui.NewDisplay(rawTerm)
	display.ClearScreen()
}
