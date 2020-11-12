package main

import (
	"machine"
	"time"
)

func main() {
	led1 := machine.D12
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led2 := machine.D13 // it's using same signal as integrated LED
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Define parameters
	var (
		min, max, wt int16 = 50, 500, 50
		step int16 = 50
	)

	for {
		if wt < min || wt > max {
			step = step * -1
		}

		wt += step

		led1.High()
		led2.Low()
		time.Sleep(time.Millisecond * time.Duration(wt))
		led1.Low()
		led2.High()
		time.Sleep(time.Millisecond * time.Duration(wt))
	}

}
