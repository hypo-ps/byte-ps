package main

import (
	"byte-ps/internal/editor"
	"fmt"
)

func main() {
	fmt.Println("hop on to byte-ps... :)")
	editor, err := editor.NewEditor()
	if err != nil {
		fmt.Printf("error while starting the editor, %v\n", err)
		return
	}

	err = editor.Begin()
	if err != nil {
		fmt.Printf("error while starting the editor, %v\n", err)
		return
	}
}
