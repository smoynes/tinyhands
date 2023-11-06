// rgbw is an experiment in RBGW driving.
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const tickInterval = 200 * time.Millisecond

var neo machine.Pin = machine.GP15
var ws ws2812.Device
var leds []color.RGBA
var whitescale [16]uint8
var colorSpace = color.RGBAModel

func init() {
	for i := 0; i < len(whitescale); i++ {
		whitescale[i] = uint8(i)
	}
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)
}

func fadeLeds(offset int) {
	c := color.NRGBA{}

	println(offset, len(whitescale))
	c.R = whitescale[offset%len(whitescale)]
	c.A = whitescale[offset%len(whitescale)]
	led := colorSpace.Convert(c).(color.RGBA)
	leds = []color.RGBA{led}
}

func main() {
	println("RGBW on", machine.Device, machine.CPUFrequency()/1000)

	flip := neo.Get()
	neo ^= neo
	neo.Set(flip)

	faderIndex := 0

	for {
		time.Sleep(tickInterval)
		//fadeLeds(faderIndex)
		//ws.WriteColors(leds)
		println(faderIndex, faderIndex%16)
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(0))
		ws.WriteByte(byte(faderIndex % 16))

		neo ^= neo
		neo.Set(flip)

		faderIndex++
	}
}
