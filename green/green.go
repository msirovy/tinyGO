package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Define parameters
	for {
		led.High()
		time.Sleep(500)
		led.Low()
		time.Sleep(5000)

	}
}
