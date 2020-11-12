package main

import (
	"machine"
	"time"
)


/*
	Use conector D9 which allows to control PWR voltage level

	~ mark on the pin means that allows to control voltage level
*/

func main() {
	// very important part!!!
	machine.InitPWM()
	led := machine.PWM{machine.D9}
	led.Configure()


	ledint := machine.LED
	ledint.Configure(machine.PinConfig{Mode: machine.PinOutput})

	min, max, voltage := 5000, 65000, 1000
	step := 5000

	//
	for {
		led.Set(uint16(voltage))

		voltage += step

		// keeps voltage between tresholds
		if voltage < min || voltage > max {
			ledint.High()
			time.Sleep(100)
			ledint.Low()
			step = step * -1
			voltage += step
		}

		time.Sleep(time.Millisecond * 50)
	}

}
