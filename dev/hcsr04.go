/*
Package dev ...

HC-SR04 is an ultrasonic distance meter
which can measure the distance to the an object like a box.

Spec:
  - power supply:	+5V DC
  - range:			2 - 450cm
  - resolution:		0.3cm
	 ___________________________
    |                           |
    |          HC-SR04          |
    |                           |
    |___________________________|
         |     |     |     |
        vcc  trig   echo  gnd

Connect to Pi:
  - vcc:	any 5v pin
  - trig:	any data pin for triggering(input)
  - echo:	any data pin for echoing(output)
  - gnd:	any gnd pin

*/
package dev

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	timeout = 3600
)

// HCSR04 ...
type HCSR04 struct {
	trig rpio.Pin
	echo rpio.Pin
}

// NewHCSR04 ...
func NewHCSR04(trig int8, echo int8) *HCSR04 {
	h := &HCSR04{
		trig: rpio.Pin(trig),
		echo: rpio.Pin(echo),
	}
	h.trig.Output()
	h.trig.Low()
	h.echo.Input()
	return h
}

// Dist is to measure the distance in cm
func (h *HCSR04) Dist() float64 {
	h.trig.Low()
	h.delay(100)
	h.trig.High()
	h.delay(15)

	for n := 0; n < timeout && h.echo.Read() != rpio.High; n++ {
		h.delay(1)
	}
	start := time.Now()

	for n := 0; n < timeout && h.echo.Read() != rpio.Low; n++ {
		h.delay(1)
	}
	return time.Now().Sub(start).Seconds() * voiceSpeed / 2.0
}

// Close ...
func (h *HCSR04) Close() {
	// do noting
}

// delay is to delay us microsecond
func (h *HCSR04) delay(us int) {
	time.Sleep(time.Duration(us) * time.Microsecond)
}
