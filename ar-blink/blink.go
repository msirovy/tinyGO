package main

import (
	"machine"
	"time"
//	"fmt"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Define parameters
	var min, max int16 = 50, 500
	var step, wt int16 = 50, 50

	for {
		if wt > max || wt < min {
			step = step * -1
		}
		wt += step

		led.High()
		time.Sleep(time.Millisecond * time.Duration(wt))
		led.Low()
		time.Sleep(time.Millisecond * time.Duration(wt))

	}
}
