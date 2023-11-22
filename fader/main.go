package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const tickInterval = 200 * time.Millisecond

var ws ws2812.Device
var leds [30]color.RGBA
var colorSpace = color.RGBAModel
var spectrum = [...]color.NRGBA{
	color.NRGBA{148, 0, 211, 0}, // Violet
	color.NRGBA{75, 0, 130, 0},  // Indigo
	color.NRGBA{0, 0, 255, 0},   // Blue
	color.NRGBA{0, 127, 255, 0}, // Blue-Green
	color.NRGBA{0, 255, 255, 0}, // Cyan
	color.NRGBA{0, 255, 127, 0}, // Aqua
	color.NRGBA{0, 255, 0, 0},   // Green
	color.NRGBA{127, 255, 0, 0}, // Lime
	color.NRGBA{255, 255, 0, 0}, // Yellow
	color.NRGBA{255, 127, 0, 0}, // Orange
	color.NRGBA{255, 0, 0, 0},   // Red
	color.NRGBA{255, 0, 127, 0}, // Pink
	color.NRGBA{255, 0, 255, 0}, // Magenta
	color.NRGBA{127, 0, 255, 0}, // Purple
}

func init() {
	fadeLeds(0)
	neo.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	ws = ws2812.New(neo)
}

func fadeLeds(offset int) {
	fade := fader[offset%len(fader)]

	for i := 0; i < len(leds); i++ {
		c := spectrum[(i+offset)%len(spectrum)]
		c.A = fade
		led := colorSpace.Convert(c)
		leds[i] = led.(color.RGBA)
	}
}

func main() {
	println("FADER on", machine.Device, machine.CPUFrequency()/1000)

	ws.WriteColors(leds[:])

	flip := neo.Get()
	neo ^= neo
	neo.Set(flip)

	faderIndex := 0

	for {
		time.Sleep(tickInterval)
		fadeLeds(faderIndex)
		ws.WriteColors(leds[:])

		neo ^= neo
		neo.Set(flip)

		faderIndex++
	}
}
