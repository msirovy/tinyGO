package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})


		min := 50
		max := 500

		step := 50


		for {
			for i := min; i < max; i=i+step {
        led.High()
        time.Sleep(time.Millisecond * time.Duration(i))
        led.Low()
        time.Sleep(time.Millisecond * time.Duration(i))
			}

			for i := max; i > min; i=i-step {
        led.High()
        time.Sleep(time.Millisecond * time.Duration(i))
        led.Low()
        time.Sleep(time.Millisecond * time.Duration(i))
			}

		}
}
