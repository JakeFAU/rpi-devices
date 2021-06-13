package main

import (
	"log"
	"time"

	"github.com/jakefau/rpi-devices/dev"
	"github.com/jakefau/rpi-devices/util"
)

func main() {
	oled, err := dev.NewOLED(128, 32)
	if err != nil {
		log.Printf("failed to create an oled, error: %v", err)
		return
	}

	util.WaitQuit(oled.Close)
	for {
		t := time.Now().Format("15:04:05")
		if err := oled.Display(t, 19, 0, 25); err != nil {
			log.Printf("failed to display time, error: %v", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
