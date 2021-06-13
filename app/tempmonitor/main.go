package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/jakefau/rpi-devices/dev"
	"github.com/jakefau/rpi-devices/iot"
	"github.com/jakefau/rpi-devices/util"
	"github.com/stianeikeland/go-rpio"
)

const (
	ledPin                 = 12
	lowTemperatureWarning  = 18
	highTemperatureWarning = 30
	intervalTime           = 1 * time.Minute
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatalf("[tempmonitor]failed to open rpio, error: %v", err)
		return
	}
	defer rpio.Close()

	temp := dev.NewDS18B20()
	if temp == nil {
		log.Printf("[tempmonitor]failed to new temperature sensor")
		return
	}
	led := dev.NewLed(ledPin)
	if led == nil {
		log.Printf("[tempmonitor]failed to new led")
		return
	}

	oneNetCfg := &iot.OneNetConfig{
		Token: iot.OneNetToken,
		API:   iot.OneNetAPI,
	}
	cloud := iot.NewCloud(oneNetCfg)
	if cloud == nil {
		log.Printf("[tempmonitor]failed to new OneNet iot cloud")
		return
	}

	monitor := tempMonitor{
		temp:  temp,
		cloud: cloud,
		led:   led,
	}

	util.WaitQuit(func() {
		rpio.Close()
	})

	monitor.start()
}

type tempMonitor struct {
	temp  *dev.DS18B20
	led   *dev.Led
	cloud iot.Cloud
}

func (m *tempMonitor) start() {
	for {
		time.Sleep(intervalTime)
		c, err := m.temp.GetTemperature()
		if err != nil {
			log.Printf("[tempmonitor]failed to get temperature, error: %v", err)
			continue
		}

		v := &iot.Value{
			Device: "temperature",
			Value:  c,
		}
		go m.cloud.Push(v)
		go m.led.Blink(5, 500)

		if c <= lowTemperatureWarning || c >= highTemperatureWarning {
			go m.notitfy(c)
		}
	}
}

func (m *tempMonitor) notitfy(temperatue float32) {
	_, err := exec.LookPath("mutt")
	if err != nil {
		log.Printf("[tempmonitor]need to install mutt for email notification")
		return
	}
	subject := "Low Temperature Warning"
	if temperatue >= highTemperatureWarning {
		subject = "High Temperature Warning"
	}
	subject = fmt.Sprintf("%v: %.2f C", subject, temperatue)
	cmd := exec.Command("mutt", "-s", subject, "youremail@xxx.com")
	if err := cmd.Run(); err != nil {
		log.Printf("[tempmonitor]failed to send email")
	}
	return
}
