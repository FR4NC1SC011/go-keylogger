package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

// TODO:
//		1. Read Keyboard
//		2. Write to file
//		3. Send to mail

func main() {
	fmt.Println("Hello world")
	readKeyboard()
}

func readKeyboard() {
	keysPressed, err := keyboard.GetKeys(10)
	check(err)

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		event := <-keysPressed
		if event.Err != nil {
			panic(event.Err)
		}
		fmt.Printf("You pressed: rune %q, key %X\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			break
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
