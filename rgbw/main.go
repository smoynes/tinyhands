// rgbw is an experiment in RBGW driving.
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const tickInterval = 40 * time.Millisecond

var ws ws2812.Device
var rgbw color.RGBA
var button = machine.BUTTONA // Button A
var lit bool

func init() {
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	button.SetInterrupt(machine.PinToggle, func(pin machine.Pin) {
		if pin.Get() != lit {
			return
		}
		if lit {
			lit = false
		} else {
			lit = true
		}
		led.Set(lit)
	})
}

func main() {
	println("RGBW on", machine.Device, machine.CPUFrequency()/1000)

	flip := neo.Get()
	neo ^= neo
	neo.Set(flip)

	var faderInc bool = true
	var faderIndex uint8 = 0

	for {
		time.Sleep(tickInterval)
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(faderIndex))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(faderIndex))

		neo ^= neo
		neo.Set(flip)

		if faderInc {
			faderIndex += 1
		} else {
			faderIndex -= 1
		}

		if faderIndex >= 24 {
			faderInc = false // dec
		} else if faderIndex <= 1 {
			faderInc = true // inc
		}
	}
}
