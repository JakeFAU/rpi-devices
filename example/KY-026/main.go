package main

import (
	"log"
	"time"

	"github.com/jakefau/rpi-devices/dev"
	"github.com/jakefau/rpi-devices/util"
	"github.com/stianeikeland/go-rpio"
)

const (
	pin = 18
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatalf("failed to open rpio, error: %v", err)
		return
	}
	defer rpio.Close()

	f := dev.NewKY026(pin)
	util.WaitQuit(func() {
		rpio.Close()
	})
	for {
		fire := f.Detected()
		if fire {
			log.Printf("Let me stand next to your fire")
		}
		time.Sleep(10 * time.Millisecond)
	}
}
