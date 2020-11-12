package main

import (
	"machine"
	"time"
)

func main() {
	led1 := machine.D12
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led2 := machine.D13
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Define parameters
	var min, max int16 = 50, 500
	var step int16 = 50

	for {
		for i := min; i < max; i=i+step {
			led1.High()
			led2.Low()
			time.Sleep(time.Millisecond * time.Duration(i))
			led1.Low()
			led2.High()
			time.Sleep(time.Millisecond * time.Duration(i))
		}

		for i := max; i > min; i=i-step {
			led1.High()
			led2.Low()
			time.Sleep(time.Millisecond * time.Duration(i))
			led1.Low()
			led1.High()
			time.Sleep(time.Millisecond * time.Duration(i))
		}

	}
}
